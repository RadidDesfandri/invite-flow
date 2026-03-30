package handler

import (
	"context"
	"net/http"
	"time"

	"backend/service"
	"backend/utils"
)

type HealthHandler struct {
	healthService *service.HealthService
}

func NewHealthHandler(healthService *service.HealthService) *HealthHandler {
	return &HealthHandler{
		healthService: healthService,
	}
}

func (h *HealthHandler) Check(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
	defer cancel()

	data, err := h.healthService.Check(ctx)
	if err != nil {
		utils.WriteError(w, http.StatusServiceUnavailable, data["status"], nil)
		return
	}

	utils.WriteSuccess(w, http.StatusOK, "health check success", data)
}
