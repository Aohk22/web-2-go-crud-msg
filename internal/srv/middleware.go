package srv

import (
	"errors"
	"log"
	"net/http"
	"strings"
)

func loggerMiddleware(logger *log.Logger, next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Printf("%v %v\n", r.Method, r.URL)

		next.ServeHTTP(w, r)
	}
}

func authMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/login" {
			cookie := r.Header.Get("Authorization")
			if len(cookie) == 0 { http.Error(w, "need jwt auth header", 401); return }
			tokens := strings.Split(cookie, " ")
			if len(tokens) < 2 { http.Error(w, "token parse error", 500); return }
			token := tokens[1]

			_, err := verifyToken(token)
			if errors.Is(err, errors.New("invalid token")) {
				http.Error(w, err.Error(), 401)
				return
			} else if err != nil {
				http.Error(w, "verify token error", 500)
				return
			}
		}

		next.ServeHTTP(w, r)
	}
}

func corsMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "http://localhost:5173")
		w.Header().Add("Access-Control-Allow-Credentials", "true")
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		w.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

	if r.Method == "OPTIONS" {
		http.Error(w, "no content", http.StatusNoContent)
		return
	}

		next.ServeHTTP(w, r);
	}
}
