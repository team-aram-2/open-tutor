package api

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"open-tutor/internal/services/db"
	"open-tutor/util"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/oapi-codegen/runtime/types"
	"golang.org/x/crypto/bcrypt"
)

func generateSessionTokenForUser(userId string, rememberLogin bool) (string, error) {
	jwtKeyPair, err := util.GetKeyPair("user-auth-jwt")
	if err != nil {
		return "", err
	}

	type sessionClaims struct {
		UserID string `json:"user_id"`
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

func applySessionTokenForUserId(userId string, rememberLogin bool, w http.ResponseWriter) error {
	sessionToken, err := generateSessionTokenForUser(userId, rememberLogin)
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
	parsedBool, err := strconv.ParseBool(r.FormValue("rememberLogin"))
	if err != nil {
		*loginData.RememberLogin = false
	} else {
		*loginData.RememberLogin = parsedBool
	}

	var (
		userId            string
		savedPasswordHash string
	)
	err = db.GetDB().QueryRow(`
		SELECT user_id, password_hash
		FROM users
		WHERE email = $1;
	`,
		loginData.Email,
	).Scan(&userId, &savedPasswordHash)
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

	err = applySessionTokenForUserId(userId, *loginData.RememberLogin, w)
	if err != nil {
		fmt.Printf("failed to apply session token: %v\n", err)
		sendError(w, http.StatusInternalServerError, "failed to apply session token")
		return
	}
	w.Header().Set("Location", r.Header.Get("Origin"))
	w.WriteHeader(http.StatusFound)
}

func (t *OpenTutor) UserRegister(w http.ResponseWriter, r *http.Request) {
	var signupData UserSignup
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

	// Generate user id //
	userId := uuid.New().String()

	// Generate password hash via bcrypt for slow hashing (safer) //
	hashBytes, err := bcrypt.GenerateFromPassword([]byte(signupData.Password), bcrypt.DefaultCost)
	if err != nil {
		sendError(w, http.StatusInternalServerError, fmt.Sprintf("error generating password hash: %v", err.Error()))
		return
	}
	passwordHash := string(hashBytes)

	_, err = db.GetDB().Exec(`
		INSERT INTO users (user_id, email, first_name, last_name, password_hash)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING user_id, email, first_name, last_name;
	`,
		userId,
		signupData.Email,
		signupData.FirstName,
		signupData.LastName,
		passwordHash,
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

	err = applySessionTokenForUserId(userId, false, w)
	if err != nil {
		fmt.Printf("failed to apply session token: %v\n", err)
		sendError(w, http.StatusInternalServerError, "failed to apply session token")
		return
	}
	w.Header().Set("Location", r.Header.Get("Origin"))
	w.WriteHeader(http.StatusFound)
}
