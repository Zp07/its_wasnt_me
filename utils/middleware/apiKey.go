package middleware

import (
	"net/http"
	"os"
)

// Middleware para verificar la API key de OpenAI
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		expectedKey := os.Getenv("APP_API_KEY")
		providedKey := r.Header.Get("X-API-Key")

		if providedKey == "" || providedKey != expectedKey {
			http.Error(w, "API key no Autorizada", http.StatusUnauthorized)
			return
		}

		// Si la API es valida, continuar con la siguiente función
		next.ServeHTTP(w, r)
	})
}
