package api

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"open-tutor/internal/services/db"
	"open-tutor/stripe_client"
	"open-tutor/util"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/oapi-codegen/runtime/types"
	"golang.org/x/crypto/bcrypt"

	"github.com/stripe/stripe-go/v81"
)

func generateSessionTokenForUser(userId string, rememberLogin bool, roleMask util.Role) (string, error) {
	jwtKeyPair, err := util.GetKeyPair("user-auth-jwt")
	if err != nil {
		return "", err
	}

	type sessionClaims struct {
		UserID   string    `json:"user_id"`
		RoleMask util.Role `json:"role_bitmask"`
		jwt.StandardClaims
	}

	var daysUntilExpiry int64
	if rememberLogin {
		daysUntilExpiry = 30
	} else {
		daysUntilExpiry = 1
	}

	claims := sessionClaims{
		userId,
		roleMask,
		jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Unix() + int64(time.Hour.Seconds())*24*daysUntilExpiry,
			Issuer:    "OpenTutor",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	tokenString, err := token.SignedString(jwtKeyPair.PrivateKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func applySessionTokenForUserId(userId string, rememberLogin bool, roleMask util.Role, w http.ResponseWriter) error {
	sessionToken, err := generateSessionTokenForUser(userId, rememberLogin, roleMask)
	if err != nil {
		return err
	}

	responseToken := fmt.Sprintf("Bearer %s", sessionToken)

	seconds_in_day := int(time.Hour.Seconds() * 24)
	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    responseToken,
		Path:     "/",
		HttpOnly: false,
		MaxAge:   30 * seconds_in_day, // 30 day expiry, max of JWT expiry lengths
		SameSite: http.SameSiteLaxMode,
	})
	return nil
}

func sendBack(w http.ResponseWriter, r *http.Request, errMsg *string) {
	redirectUrl, err := url.Parse(r.Header.Get("Origin"))
	redirectUrl.Path = "/login"
	if err != nil {
		redirectUrl = &url.URL{}
	}
	queries := redirectUrl.Query()
	if errMsg != nil {
		queries.Set("err", *errMsg)
	}
	redirectUrl.RawQuery = queries.Encode()
	w.Header().Set("Location", redirectUrl.String())
	w.WriteHeader(http.StatusFound)
}

func (t *OpenTutor) UserLogin(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		sendError(w, http.StatusInternalServerError, fmt.Sprintf("error parsing form data: %v", err))
		return
	}
	loginData := UserLogin{
		Email:         types.Email(r.FormValue("email")),
		Password:      r.FormValue("password"),
		RememberLogin: new(bool),
	}
	*loginData.RememberLogin = r.FormValue("rememberLogin") == "on"

	var (
		userId            string
		savedPasswordHash string
		roleMask          util.Role
	)
	err = db.GetDB().QueryRow(`
		SELECT user_id, password_hash, role_mask
		FROM users
		WHERE email = $1;
	`,
		loginData.Email,
	).Scan(&userId, &savedPasswordHash, &roleMask)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err := "Invalid login"
			sendBack(w, r, &err)
			return
		}

		err := err.Error()
		sendBack(w, r, &err)
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(savedPasswordHash), []byte(loginData.Password))
	if err != nil {
		var msg string
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			msg = "Invalid login"
		} else {
			msg = err.Error()
		}
		sendBack(w, r, &msg)
		return
	}

	// pass in userId and userRole bitmask
	err = applySessionTokenForUserId(userId, *loginData.RememberLogin, roleMask, w)
	if err != nil {
		fmt.Printf("failed to apply session token: %v\n", err)
		sendError(w, http.StatusInternalServerError, "failed to apply session token")
		return
	}
	w.Header().Set("Location", r.Header.Get("Origin"))
	w.WriteHeader(http.StatusFound)
}

func (t *OpenTutor) UserRegister(w http.ResponseWriter, r *http.Request) {
	var (
		signupData UserSignup
		roleMask   util.Role
	)
	err := r.ParseForm()
	if err != nil {
		sendError(w, http.StatusInternalServerError, fmt.Sprintf("error parsing form data: %v", err))
		return
	}
	firstName := r.FormValue("first_name")
	lastName := r.FormValue("last_name")
	signupData.Email = types.Email(r.FormValue("email"))
	signupData.Password = r.FormValue("password")
	signupData.FirstName = &firstName
	signupData.LastName = &lastName
	roleMask = 0

	// Generate user id //
	userId := uuid.New().String()

	// Generate password hash via bcrypt for slow hashing (safer) //
	hashBytes, err := bcrypt.GenerateFromPassword([]byte(signupData.Password), bcrypt.DefaultCost)
	if err != nil {
		sendError(w, http.StatusInternalServerError, fmt.Sprintf("error generating password hash: %v", err.Error()))
		return
	}
	passwordHash := string(hashBytes)
	roleMask = roleMask.Add(util.User)

	// Create Stripe customer for user //
	newCustomer, err := stripe_client.GetClient().Customers.New(&stripe.CustomerParams{})
	if err != nil {
		sendError(w, http.StatusInternalServerError, fmt.Sprintf("error creating Stripe customer: %v", err))
		return
	}

	// Insert into DB //
	_, err = db.GetDB().Exec(`
		INSERT INTO users (user_id, email, first_name, last_name, password_hash, role_mask, stripe_customer_id)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING user_id, email, first_name, last_name;
	`,
		userId,
		signupData.Email,
		signupData.FirstName,
		signupData.LastName,
		passwordHash,
		roleMask,
		newCustomer.ID,
	)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value") {
			msg := "An account already exists with this email"
			sendBack(w, r, &msg)
		} else {
			fmt.Printf("error creating user: %v\n", err)
			sendError(w, http.StatusInternalServerError, fmt.Sprintf("error creating user: %v", err))
		}
		return
	}

	err = applySessionTokenForUserId(userId, false, roleMask, w)
	if err != nil {
		fmt.Printf("failed to apply session token: %v\n", err)
		sendError(w, http.StatusInternalServerError, "failed to apply session token")
		return
	}
	w.Header().Set("Location", r.Header.Get("Origin"))
	w.WriteHeader(http.StatusFound)
}
