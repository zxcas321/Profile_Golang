package route

import (
	"net/http"

	"profile_go/internal/controller"
	"profile_go/internal/middleware"
	"profile_go/internal/model"
	"profile_go/pkg/helper"
)

func SetupRoutes(
	mux *http.ServeMux,
	authCtrl *controller.AuthController,
	userCtrl *controller.UserController,
) {
	// ─── Auth Routes ────────────────────────────────────────────────

	// Login (public)
	mux.HandleFunc("/api/auth/login", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			helper.WriteError(w, http.StatusMethodNotAllowed, "method not allowed")
			return
		}
		authCtrl.Login(w, r)
	})

	// Logout (protected)
	mux.HandleFunc("/api/auth/logout", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			helper.WriteError(w, http.StatusMethodNotAllowed, "method not allowed")
			return
		}
		middleware.AuthMiddleware(authCtrl.Logout)(w, r)
	})

	// Register (superuser only)
	mux.HandleFunc("/api/auth/register", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			helper.WriteError(w, http.StatusMethodNotAllowed, "method not allowed")
			return
		}
		middleware.AuthMiddleware(
			middleware.RoleMiddleware(model.RoleSuperUser, authCtrl.Register),
		)(w, r)
	})

	// ─── User Routes (protected) ────────────────────────────────────

	// GET /api/users     → GetAll
	// POST /api/users    → Create
	mux.HandleFunc("/api/users", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			middleware.AuthMiddleware(userCtrl.GetAll)(w, r)
		case http.MethodPost:
			middleware.AuthMiddleware(
				middleware.RoleMiddleware(model.RoleSuperUser, userCtrl.Create),
			)(w, r)
		default:
			helper.WriteError(w, http.StatusMethodNotAllowed, "method not allowed")
		}
	})

	// GET    /api/users/{id} → GetByID
	// PUT    /api/users/{id} → Update
	// DELETE /api/users/{id} → Delete
	mux.HandleFunc("/api/users/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			middleware.AuthMiddleware(userCtrl.GetByID)(w, r)
		case http.MethodPut:
			middleware.AuthMiddleware(userCtrl.Update)(w, r)
		case http.MethodDelete:
			middleware.AuthMiddleware(
				middleware.RoleMiddleware(model.RoleSuperUser, userCtrl.Delete),
			)(w, r)
		default:
			helper.WriteError(w, http.StatusMethodNotAllowed, "method not allowed")
		}
	})
}