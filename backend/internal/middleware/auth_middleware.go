package middleware

import (
	"context"
	"net/http"
	"strings"

	"profile_go/pkg/helper"
)

type contextKey string

const (
	UserIDKey contextKey = "user_id"
	ClaimsKey contextKey = "claims"
)

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			helper.WriteError(w, http.StatusUnauthorized, "missing or invalid token")
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := helper.ValidateToken(tokenStr)
		if err != nil {
			helper.WriteError(w, http.StatusUnauthorized, err.Error())
			return
		}

		// Store both user_id and full claims in context
		ctx := r.Context()
		ctx = context.WithValue(ctx, UserIDKey, claims.UserID)
		ctx = context.WithValue(ctx, ClaimsKey, claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}