package srv


import (
	"log"
	"fmt"
	"net/http"

	"github.com/Aohk22/web-2-go-crud-msg/internal/model"
)


func NewServer(
	logger *log.Logger,
	userStore model.UserStore,
) http.Handler {
	var handler http.Handler
	var mux *http.ServeMux = http.NewServeMux()

	addRoutes(mux, logger, userStore)

	handler = mux
	return handler
}


func addRoutes(
	mux *http.ServeMux,
	logger *log.Logger,
	userStore model.UserStore,
) {
	mux.Handle("/register", handleRegister(logger, userStore))
	mux.Handle("/login", handleLogin(logger, userStore))
	mux.Handle("/user/{id}", handleUserGet(logger, userStore))
	mux.Handle("/users", handleUserGetAll(logger, userStore))
}

func handleUserGetAll(logger *log.Logger, userStore model.UserStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := userStore.GetAllUsers()
		if err != nil { logger.Println(err) }

		for _, user := range users {
			logger.Println(user)
			fmt.Fprintf(w, "%v\n", user)
		}
	}
}


func handleRegister(logger *log.Logger, userStore model.UserStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "unimplemented\n")
	}
}


func handleUserGet(logger *log.Logger, userStore model.UserStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "unimplemented\n")
	}
}


func handleLogin(logger *log.Logger, userStore model.UserStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "unimplemented\n")
	}
}

