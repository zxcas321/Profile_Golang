package middleware

import (
	"net/http"

	"profile_go/pkg/helper"
)

func RoleMiddleware(role string, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get claims from context (set by AuthMiddleware)
		claims, ok := r.Context().Value(ClaimsKey).(*helper.Claims)
		if !ok || claims == nil {
			helper.WriteError(w, http.StatusUnauthorized, "unauthorized")
			return
		}

		if claims.Role != role {
			helper.WriteError(w, http.StatusForbidden, "forbidden: insufficient permissions")
			return
		}

		next.ServeHTTP(w, r)
	}
}