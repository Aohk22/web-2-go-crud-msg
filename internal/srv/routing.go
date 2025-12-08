package srv

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/Aohk22/web-2-go-crud-msg/internal/model"
)

func addRoutes(mux *http.ServeMux, stores *Stores) {
	mux.Handle("/register", handleRegister(stores.UserStore))
	mux.Handle("/login", handleLogin(stores.UserStore))

	mux.Handle("/users", handleUserGetAll(stores.UserStore))
	mux.Handle("/rooms", handleRoomGetAll(stores.RoomStore))

	mux.Handle("/user/{id}", handleUserGet(stores.UserStore))
	mux.Handle("/user/{id}/messages", handleMessageGet(stores.MessageStore))

	mux.Handle("/room/{id}", handleRoomGet(stores.RoomStore))

	mux.Handle("/room/{id}/users", handleRoomGetUsers(stores.UserStore))
	mux.Handle("/room/{id}/messages", handleRoomGetMessages(stores.MessageStore))
}


func handleUserGet(userStore model.UserStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := r.PathValue("id")
		id, err := strconv.ParseInt(idStr, 10, 16)
		if err != nil { http.Error(w, "cant convert id in param", 500) }
		
		user, err := userStore.GetUser(uint32(id))
		if err != nil { http.Error(w, "cant get user", 500) }
		fmt.Fprintf(w, "%v\n", user)
	}
}


func handleUserGetAll(userStore model.UserStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := userStore.GetAllUsers()
		if err != nil { http.Error(w, err.Error(), 500) }

		for _, user := range users {
			fmt.Fprintf(w, "%d - %s - %s\n", user.Id, user.Name, user.Password)
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


func handleMessageGet(messageStore model.MessageStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := r.PathValue("id")
		id, err := strconv.ParseInt(idStr, 10, 16)
		if err != nil { http.Error(w, "cant convert id in param", 500) }

		var getMsgStrat = &model.GetMessageByUser{ UserId: uint32(id) }
		msg, err := messageStore.GetMessage(getMsgStrat)
		if err != nil { http.Error(w, "cant get message", 500) }

		fmt.Fprintf(w, "%v\n", msg)
	}
}


func handleRoomGet(roomStore model.RoomStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := r.PathValue("id")
		id, err := strconv.ParseInt(idStr, 10, 16)
		if err != nil { http.Error(w, "cant convert id in param", 500) }

		room, err := roomStore.GetRoom(uint32(id))
		if err != nil { http.Error(w, "can get room", 500) }

		fmt.Fprintf(w, "%v\n", room)
	}
}


func handleRoomGetAll(roomStore model.RoomStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rooms, err := roomStore.GetAllRooms()
		if err != nil { http.Error(w, "cannot get all rooms", 500) }

		for _, room := range rooms {
			fmt.Fprintf(w, "Room name: %s\nCreation date: %v\n\n", room.Name, room.Time)
		}
	}
}


func handleRoomGetUsers(userStore model.UserStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		roomIdStr := r.PathValue("id")
		roomIdInt, err := strconv.ParseInt(roomIdStr, 10, 16)
		if err != nil { http.Error(w, "could not parse int.", 500) }

		users, err := userStore.GetUsersByRoom(uint32(roomIdInt))
		if err != nil { http.Error(w, "could not get users by room", 500) }

		for _, user := range users {
			fmt.Fprintf(w, "Room %d: %v\n", roomIdInt, user)
		}
	}
}


func handleRoomGetMessages(messageStore model.MessageStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		roomIdStr := r.PathValue("id")
		roomIdInt, err := strconv.ParseInt(roomIdStr, 10, 16)

		strat := &model.GetMessageByRoom { RoomId: uint32(roomIdInt)}
		messages, err := messageStore.GetMessage(strat)
		if err != nil { http.Error(w, "cannot get message by room", 500) }

		for _, message := range messages {
			fmt.Fprintf(w, "message: %v\n", message)
		}
	}
}
