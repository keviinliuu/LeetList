package auth

import (
	"context"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

type contextKey struct {
	name string
}

var UserCtxKey = &contextKey{"user"}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")

		if header == "" {
			next.ServeHTTP(w, r)
			return
		}

		splits := strings.Split(header, " ")
		if len(splits) != 2 {
			next.ServeHTTP(w, r)
			return
		}

		tokenStr := splits[1]
		claims := &jwt.MapClaims{}
		token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil || !token.Valid {
			next.ServeHTTP(w, r)
			return 
		}

		email := (*claims)["user_id"]

		ctx := context.WithValue(r.Context(), UserCtxKey, email)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}