package middleware

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	"open-tutor/internal/services/db"
	"open-tutor/util"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

const (
	AuthenticationContextKey util.HttpContextKey = "user_auth"
)

type Claims struct {
	UserID   *string       `json:"user_id"`
	RoleMask util.RoleMask `json:"role_mask"`
	jwt.StandardClaims
}

type AuthenticationInfo struct {
	UserID   string
	RoleMask util.RoleMask
}

func GetAuthenticationInfo(r *http.Request) *AuthenticationInfo {
	authInfo, ok := r.Context().Value(AuthenticationContextKey).(AuthenticationInfo)
	if !ok {
		return nil
	}
	return &authInfo
}

func InvalidateAuthRedirect(r *http.Request, w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    "",
		Path:     "/",
		HttpOnly: false,
		MaxAge:   -1, // Expire immediately
	})
	// Redirect the user to the login page
	http.Redirect(w, r, "/login", http.StatusForbidden)
}

func Authenticate(next http.Handler) http.HandlerFunc {
	jwtKeyPair, err := util.GetKeyPair("user-auth-jwt")
	if err != nil {
		panic(fmt.Sprintf("failed to get user-auth-jwt keypair: %v", err))
	}

	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("session_token")
		if err != nil {
			// No session_token cookie found //
			fmt.Printf("[Authenticate] error reading session_token cookie: %v\n", err)
			next.ServeHTTP(w, r)
			return
		}

		// Parse session_token //
		sessionToken := strings.TrimPrefix(cookie.Value, "Bearer ")
		token, err := jwt.ParseWithClaims(sessionToken, &Claims{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return jwtKeyPair.PublicKey, nil
		})
		if err != nil {
			log.Printf("[Authenticate] invalid token: %v\n", err)
			InvalidateAuthRedirect(r, w)
			return
		}

		claims, ok := token.Claims.(*Claims)
		if !ok || !token.Valid {
			log.Printf("[Authenticate] invalid token claims: %v\n", err)
			InvalidateAuthRedirect(r, w)
			return
		}

		// Parse user ID safely
		var userId string
		if claims.UserID != nil {
			parsedId, err := uuid.Parse(*claims.UserID)
			if err == nil {
				userId = parsedId.String()
			} else {
				log.Printf("[Authenticate] failed to parse uuid from user id: %v\n", err)
				InvalidateAuthRedirect(r, w)
				return
			}
		}

		// Get the current role from the database
		var currentRoleMask util.RoleMask
		err = db.GetDB().QueryRow("SELECT role_mask FROM users WHERE user_id = $1", userId).Scan(&currentRoleMask)
		if err != nil {
			// Handle error if the user is not found or DB issue
			log.Printf("[Authenticate] failed to find user in database: %v\n", err)
			InvalidateAuthRedirect(r, w)
			return
		}

		// Compare the stored role with the current role
		if claims.RoleMask != currentRoleMask {
			log.Printf("[Authenticate] JWT.Claims role bitmask (%v)!= databse role bitmask (%v).", claims.RoleMask, currentRoleMask)
			InvalidateAuthRedirect(r, w)
			return
		}

		// Add the userId to authentication context //
		authInfo := AuthenticationInfo{
			UserID:   userId,
			RoleMask: claims.RoleMask,
		}
		ctx := context.WithValue(r.Context(), AuthenticationContextKey, authInfo)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}
