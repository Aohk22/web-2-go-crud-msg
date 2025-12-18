package srv

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/Aohk22/web-2-go-crud-msg/internal/model"
)


func addRoutes(ctx context.Context, mux *http.ServeMux, stores *Stores) {
	mux.Handle("/register", handleRegister(ctx, stores.UserStore))
	mux.Handle("/login", handleLogin(ctx, stores.UserStore))

	mux.Handle("/users", handleUserGetAll(ctx, stores.UserStore))
	mux.Handle("/rooms", handleRoomGetAll(ctx, stores.RoomStore))

	mux.Handle("/user/{id}", handleUserGet(ctx, stores.UserStore))
	mux.Handle("/user/{id}/messages", handleMessageGet(ctx, stores.MessageStore))

	mux.Handle("/room/{id}", handleRoomGet(ctx, stores.RoomStore))

	mux.Handle("/room/{id}/users", handleRoomGetUsers(ctx, stores.UserStore))
	mux.Handle("/room/{id}/messages", handleRoomGetMessages(ctx, stores.MessageStore))
}


func handleUserGet(ctx context.Context, userStore model.UserStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := r.PathValue("id")
		id, err := strconv.ParseInt(idStr, 10, 16)
		if err != nil { http.Error(w, "cant convert id in param", 500); return }
		
		user, err := userStore.GetUser(ctx, uint32(id))
		if err != nil { http.Error(w, "cant get user", 500); return }

		json, err := json.Marshal(user)
		if err != nil { http.Error(w, "cant convert json", 500); return }

		fmt.Fprintf(w, "%s", json)
	}
}


func handleUserGetAll(ctx context.Context, userStore model.UserStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := userStore.GetAllUsers(ctx)
		if err != nil { http.Error(w, err.Error(), 500); return }

		json, err := json.Marshal(users)
		if err != nil { http.Error(w, "cant convert json", 500); return }

		fmt.Fprintf(w, "%s", json)
	}
}


func handleRegister(ctx context.Context, userStore model.UserStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			var user model.User

			data, err := io.ReadAll(r.Body)
			if err != nil { http.Error(w, "error reading json", 500); return }
			if validJson := json.Valid(data); !validJson { http.Error(w, "not valid json", 500); return }

			if err := json.Unmarshal(data, &user); err != nil {
				http.Error(w, "error parsing json", 500); return
			}

			stat, err := userStore.AddUser(ctx, user.Name, user.Password)
			if err != nil { http.Error(w, "add user error", 500); return }
			fmt.Fprintf(w, "%s", stat)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}


func handleLogin(ctx context.Context, userStore model.UserStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		content, err := io.ReadAll(r.Body)
		if err != nil { http.Error(w, "failed to read json data", 500); return }
		defer r.Body.Close()

		creds := struct {
			Name string `json:"username"`
			Password string `json:"password"`
		}{}

		err = json.Unmarshal(content, &creds)
		if err != nil { http.Error(w, "unmarshal error", 500); return }

		valid, err := userStore.CheckUser(ctx, creds.Name, creds.Password)
		if err != nil { http.Error(w, "checkuser error", 500); return }
		user, err := userStore.GetUserId(ctx, creds.Name)
		if err != nil { http.Error(w, "get user id error", 500); return }
		userIdStr := strconv.Itoa(int(user.Id))
		if !valid { 
			http.Error(w, "invalid user", 401) 
			return
		} else {
			tokenString, err := createToken(creds.Name, userIdStr)
			if err != nil { http.Error(w, "could not create jwt", 500); return }
			fmt.Fprintf(w, "%s", tokenString)
		}
	}
}


func handleMessageGet(ctx context.Context, messageStore model.MessageStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := r.PathValue("id")
		id, err := strconv.ParseInt(idStr, 10, 16)
		if err != nil { http.Error(w, "cant convert id in param", 500); return }

		var getMsgStrat = &model.GetMessageByUser{ UserId: uint32(id) }
		msg, err := messageStore.GetMessage(ctx, getMsgStrat)
		if err != nil { http.Error(w, "cant get message", 500); return }

		json, err := json.Marshal(msg)
		if err != nil { http.Error(w, "cant convert json", 500); return }

		fmt.Fprintf(w, "%s", json)
	}
}


func handleRoomGet(ctx context.Context, roomStore model.RoomStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := r.PathValue("id")
		id, err := strconv.ParseInt(idStr, 10, 16)
		if err != nil { http.Error(w, "cant convert id in param", 500); return }

		room, err := roomStore.GetRoom(ctx, uint32(id))
		if err != nil { http.Error(w, "cant get room", 500); return }

		json, err := json.Marshal(room)
		if err != nil { http.Error(w, "cant convert json", 500); return }

		fmt.Fprintf(w, "%s", json)
	}
}


func handleRoomGetAll(ctx context.Context, roomStore model.RoomStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rooms, err := roomStore.GetAllRooms(ctx)
		if err != nil { http.Error(w, "cannot get all rooms", 500); return }

		json, err := json.Marshal(rooms)
		if err != nil { http.Error(w, "cant convert json", 500); return }

		fmt.Fprintf(w, "%s", json)
	}
}


func handleRoomGetUsers(ctx context.Context, userStore model.UserStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		roomIdStr := r.PathValue("id")
		roomIdInt, err := strconv.ParseInt(roomIdStr, 10, 16)
		if err != nil { http.Error(w, "could not parse int.", 500); return }

		users, err := userStore.GetUsersByRoom(ctx, uint32(roomIdInt))
		if err != nil { http.Error(w, "could not get users by room", 500); return }

		json, err := json.Marshal(users)
		if err != nil { http.Error(w, "cant convert json", 500); return }

		fmt.Fprintf(w, "%s", json)
	}
}


func handleRoomGetMessages(ctx context.Context, messageStore model.MessageStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		roomIdStr := r.PathValue("id")
		roomIdInt, err := strconv.ParseInt(roomIdStr, 10, 16)

		strat := &model.GetMessageByRoom { RoomId: uint32(roomIdInt)}
		messages, err := messageStore.GetMessage(ctx, strat)
		if err != nil { http.Error(w, "cannot get message by room", 500); return }

		json, err := json.Marshal(messages)
		if err != nil { http.Error(w, "cant convert json", 500); return }

		fmt.Fprintf(w, "%s", json)
	}
}
