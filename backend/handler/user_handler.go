package handler

import (
	"errors"
	"net/http"
	"strconv"

	"backend/service"
	"backend/utils"
)

type UserHandler struct {
	userService *service.UserService
}

type createUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req createUserRequest
	if err := utils.DecodeJSON(r, &req); err != nil {
		utils.WriteError(w, http.StatusBadRequest, "invalid request payload", nil)
		return
	}

	user, err := h.userService.Create(r.Context(), req.Name, req.Email)
	if err != nil {
		if errors.Is(err, service.ErrInvalidUserInput) {
			utils.WriteError(w, http.StatusBadRequest, err.Error(), nil)
			return
		}
		utils.WriteError(w, http.StatusInternalServerError, "failed to create user", nil)
		return
	}

	utils.WriteSuccess(w, http.StatusCreated, "user created", user)
}

func (h *UserHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	idValue := r.PathValue("id")
	id, err := strconv.ParseInt(idValue, 10, 64)
	if err != nil || id <= 0 {
		utils.WriteError(w, http.StatusBadRequest, "invalid user id", nil)
		return
	}

	user, err := h.userService.GetByID(r.Context(), id)
	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			utils.WriteError(w, http.StatusNotFound, "user not found", nil)
			return
		}
		utils.WriteError(w, http.StatusInternalServerError, "failed to get user", nil)
		return
	}

	utils.WriteSuccess(w, http.StatusOK, "user found", user)
}

func (h *UserHandler) List(w http.ResponseWriter, r *http.Request) {
	limit := parseIntWithDefault(r.URL.Query().Get("limit"), 10)
	offset := parseIntWithDefault(r.URL.Query().Get("offset"), 0)

	users, err := h.userService.List(r.Context(), limit, offset)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, "failed to list users", nil)
		return
	}

	utils.WriteSuccess(w, http.StatusOK, "users retrieved", map[string]any{
		"items":  users,
		"limit":  limit,
		"offset": offset,
	})
}

func parseIntWithDefault(value string, fallback int) int {
	if value == "" {
		return fallback
	}
	number, err := strconv.Atoi(value)
	if err != nil {
		return fallback
	}
	return number
}
