package middlewares

import (
	"context"
	s "resume-analyzer/internal/shared"
	"resume-analyzer/pkg/jwt"
	"net/http"
	"strings"
)

func Auth(Jwtsecret string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				s.ReqResponse(w, http.StatusUnauthorized, s.Payload{Message: "no token found"})
				return
			}
			tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

			claims, err := jwt.VerifyToken(tokenStr, Jwtsecret)
			if err != nil {
				s.ReqResponse(w, http.StatusUnauthorized, s.Payload{Message: "invalid token"})
				return
			}

			ctx := context.WithValue(r.Context(), userCtxKey, claims)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func AuthAdmin(Jwtsecret string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				s.ReqResponse(w, http.StatusUnauthorized, s.Payload{Message: "no token found"})
				return
			}
			tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

			claims, err := jwt.VerifyToken(tokenStr, Jwtsecret)
			if err != nil {
				s.ReqResponse(w, http.StatusUnauthorized, s.Payload{Message: "invalid token"})
				return
			}

			if claims.Role != "admin" {
				s.ReqResponse(w, http.StatusUnauthorized, s.Payload{Message: "Unauthorized user"})
				return
			}

			ctx := context.WithValue(r.Context(), userCtxKey, claims)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
