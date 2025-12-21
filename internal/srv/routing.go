package srv

import (
	"context"
	"net/http"
)

func addRoutes(ctx context.Context, mux *http.ServeMux, stor *Stores) {
	mux.Handle("POST /register", register(ctx, stor.UserStore))
	mux.Handle("POST /login", login(ctx, stor.UserStore))

	mux.Handle("GET /users", getUsers(ctx, stor.UserStore))
	mux.Handle("GET /user/{id}", getUser(ctx, stor.UserStore))
	mux.Handle("GET /user/{id}/messages", getMessage(ctx, stor.MessageStore))

	mux.Handle("GET /rooms", getRooms(ctx, stor.RoomStore))
	mux.Handle("GET /room/{id}", getRoom(ctx, stor.RoomStore))
	mux.Handle("GET /room/{id}/users", getRoomUsers(ctx, stor.UserStore))
	mux.Handle("GET /room/{id}/messages", getRoomMessages(ctx, stor.MessageStore))

	mux.Handle("POST /room/{id}/messages", getOldMessages(ctx, stor.MessageStore))

	mux.Handle("PUT /room/{id}", putRoom(ctx, stor))
}

