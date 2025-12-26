package v1

import (
	"resume-analyzer/internal/handlers"
	"resume-analyzer/internal/routes/v1/auth"

	"github.com/gorilla/mux"
)

func NewV1Router(ah *handlers.AuthHandler) *mux.Router {
	r := mux.NewRouter()

	api := r.PathPrefix("/api/v1").Subrouter()

	auth.RegisterAuthRouter(api, ah)

	return r
}