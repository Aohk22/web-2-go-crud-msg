package srv

import (
	"log"
	"net/http"
)


func loggerMiddleware(logger *log.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			logger.Printf("%v\n", r.URL)
			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	}
}
