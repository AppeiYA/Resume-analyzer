package auth

import (
	"log"

	"github.com/gorilla/mux"
)

func RegisterAuthRouter(r *mux.Router) {
	auth := r.PathPrefix("/auth").Subrouter()
	log.Println(auth)
}