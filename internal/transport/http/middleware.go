package http

import (
	"context"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"log/slog"
	"net/http"
	"os"
)

func (h *Handler) IsAdminMW(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check if the user is an admin
		tokenString := r.Header.Get("Authorization")
		if len(tokenString) < 7 || tokenString[:7] != "Bearer " {
			slog.Error("invalid token", "token", tokenString)
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}
		tokenString = tokenString[7:]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				slog.Error("unexpected signing method", "method", token.Header["alg"])
				return nil, errors.New("unexpected signing method")
			}
			return []byte(os.Getenv("JWT_SECRET")), nil
		})
		if err != nil || !token.Valid {
			slog.Error("invalid token", "error", err)
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			slog.Error("invalid claims")
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}
		if claims["role"] != "admin" {
			slog.Error("user is not an admin")
			http.Error(w, "user is not an admin", http.StatusForbidden)
			return
		}
		ctx := context.WithValue(r.Context(), "userID", claims["userID"])
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
