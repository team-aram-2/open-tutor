package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"open-tutor/util"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

const (
	AuthenticationContextKey util.HttpContextKey = "user_auth"
)

type Claims struct {
	UserID *string `json:"user_id"`
	jwt.StandardClaims
}

type AuthenticationInfo struct {
	UserID string
}

func GetAuthenticationInfo(r *http.Request) *AuthenticationInfo {
	authInfo, ok := r.Context().Value(AuthenticationContextKey).(AuthenticationInfo)
	if !ok {
		return nil
	}
	return &authInfo
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
			fmt.Printf("[Authenticate] invalid token: %v\n", err)
			next.ServeHTTP(w, r)
			return
		}

		claims, ok := token.Claims.(*Claims)
		if !ok || !token.Valid {
			fmt.Printf("[Authenticate] invalid token claims: %v\n", err)
			next.ServeHTTP(w, r)
			return
		}
		var userId uuid.UUID
		if claims.UserID != nil {
			userId, err = uuid.Parse(*claims.UserID)
		}
		if err != nil {
			fmt.Printf("[Authenticate] failed to parse uuid from user id: %v\n", err)
			next.ServeHTTP(w, r)
			return
		}

		// Add the userId to authentication context //
		authInfo := AuthenticationInfo{
			UserID: userId.String(),
		}
		ctx := context.WithValue(r.Context(), AuthenticationContextKey, authInfo)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}
