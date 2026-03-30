package handler

import (
	"net/http"
	"time"

	"backend/utils"
)

type SystemHandler struct{}

func NewSystemHandler() *SystemHandler {
	return &SystemHandler{}
}

func (h *SystemHandler) Root(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		h.NotFound(w, r)
		return
	}

	utils.WriteSuccess(w, http.StatusOK, "backend is running", map[string]string{
		"service": "backend",
	})
}

func (h *SystemHandler) Ping(w http.ResponseWriter, r *http.Request) {
	utils.WriteSuccess(w, http.StatusOK, "pong", map[string]any{
		"message": "pong",
		"time":    time.Now().UTC().Format(time.RFC3339),
	})
}

func (h *SystemHandler) NotFound(w http.ResponseWriter, r *http.Request) {
	utils.WriteError(w, http.StatusNotFound, "route not found", nil)
}
