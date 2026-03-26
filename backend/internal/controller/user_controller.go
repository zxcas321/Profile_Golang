package controller

import (
	"encoding/json"
	"net/http"
	"strings"

	"profile_go/internal/request"
	"profile_go/internal/service"
	"profile_go/pkg/helper"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController(userService *service.UserService) *UserController {
	return &UserController{userService: userService}
}

func (c *UserController) Create(w http.ResponseWriter, r *http.Request) {
	var req request.CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helper.WriteError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if req.Email == "" || req.Password == "" {
		helper.WriteError(w, http.StatusBadRequest, "email and password are required")
		return
	}

	if err := c.userService.Create(req); err != nil {
		helper.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	helper.WriteSuccess(w, http.StatusCreated, "user created successfully")
}

func (c *UserController) GetAll(w http.ResponseWriter, r *http.Request) {
	users, err := c.userService.GetAll()
	if err != nil {
		helper.WriteError(w, http.StatusInternalServerError, "failed to fetch users")
		return
	}

	helper.WriteJSON(w, http.StatusOK, users)
}

func (c *UserController) GetByID(w http.ResponseWriter, r *http.Request) {
	id := extractID(r.URL.Path, "/api/users/")
	if id == "" {
		helper.WriteError(w, http.StatusBadRequest, "missing user id")
		return
	}

	user, err := c.userService.GetByID(id)
	if err != nil {
		helper.WriteError(w, http.StatusNotFound, err.Error())
		return
	}

	helper.WriteJSON(w, http.StatusOK, user)
}

func (c *UserController) Update(w http.ResponseWriter, r *http.Request) {
	id := extractID(r.URL.Path, "/api/users/")
	if id == "" {
		helper.WriteError(w, http.StatusBadRequest, "missing user id")
		return
	}

	var req request.UpdateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helper.WriteError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := c.userService.Update(id, req); err != nil {
		helper.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	helper.WriteSuccess(w, http.StatusOK, "user updated successfully")
}


func (c *UserController) Delete(w http.ResponseWriter, r *http.Request) {
	id := extractID(r.URL.Path, "/api/users/")
	if id == "" {
		helper.WriteError(w, http.StatusBadRequest, "missing user id")
		return
	}

	if err := c.userService.Delete(id); err != nil {
		helper.WriteError(w, http.StatusNotFound, err.Error())
		return
	}

	helper.WriteSuccess(w, http.StatusOK, "user deleted successfully")
}

// helper to extract ID from URL path
func extractID(path, prefix string) string {
	return strings.TrimPrefix(path, prefix)
}