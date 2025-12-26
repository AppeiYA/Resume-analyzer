package auth

import (
	"resume-analyzer/internal/handlers"

	"github.com/gorilla/mux"
)

func RegisterAuthRouter(r *mux.Router, h *handlers.AuthHandler) {
	auth := r.PathPrefix("/auth").Subrouter()
	auth.HandleFunc("/register", h.CreateUser)
}