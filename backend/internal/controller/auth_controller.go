package controller

import (
	"encoding/json"
	"net/http"

	"profile_go/internal/request"
	"profile_go/internal/service"
	"profile_go/pkg/helper"
)

type AuthController struct {
	authService *service.AuthService
}

func NewAuthController(authService *service.AuthService) *AuthController {
	return &AuthController{authService: authService}
}

func (c *AuthController) Register(w http.ResponseWriter, r *http.Request) {
	var req request.RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helper.WriteError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if req.Email == "" || req.Password == "" {
		helper.WriteError(w, http.StatusBadRequest, "email and password are required")
		return
	}

	if err := c.authService.Register(req); err != nil {
		helper.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	helper.WriteSuccess(w, http.StatusCreated, "registration successful")
}

func (c *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	var req request.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helper.WriteError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if req.Email == "" || req.Password == "" {
		helper.WriteError(w, http.StatusBadRequest, "email and password are required")
		return
	}

	token, err := c.authService.Login(req)
	if err != nil {
		helper.WriteError(w, http.StatusUnauthorized, err.Error())
		return
	}

	helper.WriteJSON(w, http.StatusOK, map[string]string{"token": token})
}

func (c *AuthController) Logout(w http.ResponseWriter, r *http.Request) {
	// JWT is stateless — client deletes the token
	// Optionally add token blacklist with Redis later
	helper.WriteSuccess(w, http.StatusOK, "logout successful")
}