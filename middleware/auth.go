package middleware

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
)

// Middleware untuk user
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Mengambil token dari header 'token'
		authHeader := r.Header.Get("token")
		if authHeader == "" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]interface{}{"status": false, "message": "Unauthorized", "errors": []string{"Token is required"}})
			return
		}

		// Cek validasi token
		token, err := jwt.Parse(authHeader, func(token *jwt.Token) (interface{}, error) {
			return []byte("secret_key"), nil
		})

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// Simpan id user ke context agar bisa digunakan di handler
			ctx := context.WithValue(r.Context(), "user_id", claims["id"])
			ctx = context.WithValue(ctx, "is_admin", claims["is_admin"])
			next.ServeHTTP(w, r.WithContext(ctx))
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]interface{}{"status": false, "message": "Unauthorized", "errors": []string{err.Error()}})
		}
	})
}

// Middleware untuk admin
func AdminMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Mengambil data isAdmin dari context token
		isAdmin := r.Context().Value("is_admin").(bool)

		if !isAdmin {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status":  false,
				"message": "Forbidden: Hanya Admin yang dapat mengakses ini",
			})
			return
		}
		next.ServeHTTP(w, r)
	})
}
