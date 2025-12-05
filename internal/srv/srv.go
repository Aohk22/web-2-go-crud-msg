package srv


import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/Aohk22/web-2-go-crud-msg/internal/model"
)


func NewServer(
	logger *log.Logger,
	userStore model.UserStore,
	messageStore model.MessageStore,
	roomStore model.RoomStore,
) http.Handler {
	var handler http.Handler
	var mux *http.ServeMux = http.NewServeMux()

	addRoutes(mux, logger, userStore, messageStore, roomStore)

	// TODO: wrap the loggers and auth middleware
	handler = mux

	return handler
}


func addRoutes(
	mux *http.ServeMux,
	logger *log.Logger,
	userStore model.UserStore,
	messageStore model.MessageStore,
	roomStore model.RoomStore,
) {
	mux.Handle("/register", handleRegister(userStore))
	mux.Handle("/login", handleLogin(userStore))
	mux.Handle("/user/{id}", handleUserGet(userStore))
	mux.Handle("/users", handleUserGetAll(userStore))

	mux.Handle("/message/{id}", handleMessageGet(messageStore))

	mux.Handle("/room/{id}", handleRoomGet(roomStore))
}


// --- user handlers ---


func handleUserGet(userStore model.UserStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := r.PathValue("id")
		id, err := strconv.ParseInt(idStr, 10, 16)
		if err != nil { http.Error(w, "cant convert id in param", 500) }
		
		user, err := userStore.GetUser(uint16(id))
		if err != nil { http.Error(w, "cant get user", 500) }
		fmt.Fprintf(w, "%v\n", user)
	}
}


func handleUserGetAll(userStore model.UserStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := userStore.GetAllUsers()
		if err != nil { http.Error(w, err.Error(), 500) }

		for _, user := range users {
			fmt.Fprintf(w, "%v\n", user)
		}
	}
}


func handleRegister(userStore model.UserStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			data, err := io.ReadAll(r.Body)
			if err != nil { http.Error(w, "error reading json", 500) }
			if validJson := json.Valid(data); !validJson { http.Error(w, "not valid json", 500) }
			var user model.User
			if err := json.Unmarshal(data, &user); err != nil {
				http.Error(w, "error parsing json", 500)
			}
			stat, err := userStore.AddUser(user.Name, user.Password)
			if err != nil { http.Error(w, "add user error", 500) }
			fmt.Fprintf(w, "%s\n", stat)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}


func handleLogin(userStore model.UserStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "unimplemented\n")
	}
}


// --- messages handlers ---


func handleMessageGet(messageStore model.MessageStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := r.PathValue("id")
		id, err := strconv.ParseInt(idStr, 10, 16)
		if err != nil { http.Error(w, "cant convert id in param", 500) }

		var getMsgStrat = &model.GetMessageByUser{ UserId: uint16(id) }
		msg, err := messageStore.GetMessage(getMsgStrat)
		if err != nil { http.Error(w, "cant get message", 500) }

		fmt.Fprintf(w, "%v\n", msg)
	}
}
