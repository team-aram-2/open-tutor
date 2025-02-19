package api

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"open-tutor/internal/services/db"
	"open-tutor/util"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/oapi-codegen/runtime/types"
	"golang.org/x/crypto/bcrypt"
)

func generateSessionTokenForUser(userId string, roleMask uint16) (string, error) {
	jwtKeyPair, err := util.GetKeyPair("user-auth-jwt")
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256,
		jwt.MapClaims{
			"user_id": userId,
			"roles": roleMask,
		})

	tokenString, err := token.SignedString(jwtKeyPair.PrivateKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func applySessionTokenForUserId(userId string, roleMask uint16, w http.ResponseWriter) error {
	sessionToken, err := generateSessionTokenForUser(userId, roleMask)
	if err != nil {
		return err
	}

	responseToken := fmt.Sprintf("Bearer %s", sessionToken)
	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    responseToken,
		Path:     "/",
		HttpOnly: true,
		MaxAge:   86400 * 1, // 1 day expiry
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
	var loginData UserLogin
	err := r.ParseForm()
	if err != nil {
		sendError(w, http.StatusInternalServerError, fmt.Sprintf("error parsing form data: %v", err))
		return
	}
	loginData.Email = types.Email(r.FormValue("email"))
	loginData.Password = r.FormValue("password")

	var (
		userId            	string
		savedPasswordHash 	string
		roleMask			uint16
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
	err = applySessionTokenForUserId(userId, ,role_mask, w)
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

	err = applySessionTokenForUserId(userId, w)
	if err != nil {
		fmt.Printf("failed to apply session token: %v\n", err)
		sendError(w, http.StatusInternalServerError, "failed to apply session token")
		return
	}
	w.Header().Set("Location", r.Header.Get("Origin"))
	w.WriteHeader(http.StatusFound)
}
