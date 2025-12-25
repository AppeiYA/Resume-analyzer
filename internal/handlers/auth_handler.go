package handlers

import (
	"encoding/json"
	"net/http"
	"resume-analyzer/internal/errors/apperrors"
	"resume-analyzer/internal/models"
	"resume-analyzer/internal/services"
	s "resume-analyzer/internal/shared"
	"resume-analyzer/internal/utils"
)

type AuthHandler struct {
	authService *services.AuthService
}

func NewAuthServiceHandler(authService *services.AuthService) *AuthHandler {
	return  &AuthHandler{authService: authService}
}

func (h *AuthHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var payload models.CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		s.ReqResponse(w, http.StatusBadRequest, s.Payload{Message: "No body in request"})
		return
	}
	if err := validate.Struct(payload); err != nil {
		errs := utils.ValidationErrors(err)
		s.ReqResponse(w, http.StatusBadRequest, s.Payload{Message: "validation error", Errors: errs})
		return
	}

	err := h.authService.CreateUser(ctx, &payload)
	if err != nil {
		if appErr, ok := err.(*apperrors.ErrorResponse); ok {
			s.ReqResponse(w, appErr.StatusCode, s.Payload{Message: appErr.Message})
			return
		}

		s.ReqResponse(w, http.StatusInternalServerError, s.Payload{Message: "internal server error"})
		return
	}

	s.ReqResponse(w, http.StatusAccepted, s.Payload{
		Message: "User Signup successful",
	})
}