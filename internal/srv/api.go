package srv

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/Aohk22/web-2-go-crud-msg/internal/model"
)

type RoomRequestData struct {
	DataType string `json:"dataType"`
	Data struct {
		Content string `json:"content"`
		Rid string `json:"rid"`
		Uid string `json:"uid"`
	} `json:"data"`
}

func getUser(ctx context.Context, userStore model.UserStore) http.HandlerFunc {
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

func getUsers(ctx context.Context, userStore model.UserStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := userStore.GetAllUsers(ctx)
		if err != nil { http.Error(w, err.Error(), 500); return }

		json, err := json.Marshal(users)
		if err != nil { http.Error(w, "cant convert json", 500); return }

		fmt.Fprintf(w, "%s", json)
	}
}

func register(ctx context.Context, userStore model.UserStore) http.HandlerFunc {
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

func login(ctx context.Context, userStore model.UserStore) http.HandlerFunc {
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

func getMessage(ctx context.Context, messageStore model.MessageStore) http.HandlerFunc {
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

func getRoom(ctx context.Context, rs model.RoomStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		roomId := r.PathValue("id")
		roomIdInt, err := strconv.ParseInt(roomId, 10, 32)
		if err != nil { http.Error(w, "cant convert id in param", 500); return }

		room, err := rs.GetRoom(ctx, uint32(roomIdInt))
		if err != nil { http.Error(w, "cant get room", 500); return }

		json, err := json.Marshal(room)
		if err != nil { http.Error(w, "cant convert to json", 500); return }

		fmt.Fprintf(w, "%s", json)
	}
}

func getRooms(ctx context.Context, roomStore model.RoomStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rooms, err := roomStore.GetAllRooms(ctx)
		if err != nil { http.Error(w, "cannot get all rooms", 500); return }

		json, err := json.Marshal(rooms)
		if err != nil { http.Error(w, "cant convert json", 500); return }

		fmt.Fprintf(w, "%s", json)
	}
}

func getRoomUsers(ctx context.Context, userStore model.UserStore) http.HandlerFunc {
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

func getRoomMessages(ctx context.Context, messageStore model.MessageStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		roomIdStr := r.PathValue("id")
		roomIdInt, err := strconv.ParseInt(roomIdStr, 10, 16)
		if err != nil { http.Error(w, "cant parse int", 500); return }

		strat := &model.GetMessageByRoom { RoomId: uint32(roomIdInt), Time: time.Now() }
		messages, err := messageStore.GetMessage(ctx, strat)
		if err != nil { 
			fmt.Println(err)
			http.Error(w, "cannot get message by room", 500); return }

		json, err := json.Marshal(messages)
		if err != nil { http.Error(w, "cant convert json", 500); return }

		fmt.Fprintf(w, "%s", json)
	}
}

func getOldMessages(ctx context.Context, messageStore model.MessageStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		roomIdStr := r.PathValue("id")
		roomIdInt, err := strconv.ParseInt(roomIdStr, 10, 16)
		if err != nil { http.Error(w, "cant parse int", 500); return }

		body := r.Body
		defer body.Close()

		data, err := io.ReadAll(body)
		if err != nil { http.Error(w, "cant read body", 500); return }

		var dataJson struct { Time string `json:"time"` }
		err = json.Unmarshal(data, &dataJson)
		if err != nil { http.Error(w, "cant unmarshal data", 500); return }
		ts, err := strconv.ParseInt(dataJson.Time, 10, 64)
		if err != nil { http.Error(w, "invalid time" , 500); return }
		if ts > 1e12 {
			ts = ts/1000
		}

		strat := &model.GetMessageByRoom { RoomId: uint32(roomIdInt), Time: time.Unix(ts, 0).UTC() }
		messages, err := messageStore.GetMessage(ctx, strat)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "failed to get message", 500)
			return
		}

		json, err := json.Marshal(messages)
		if err != nil { http.Error(w, "cant convert json", 500); return }

		fmt.Fprintf(w, "%s", json)
	}
}

func putRoom(ctx context.Context, s *Stores) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Content-Type") != "application/json" {
			http.Error(w, "wrong content type", 500)
			return
		}
		requestData := RoomRequestData{}
		body := r.Body
		defer body.Close()
		data, err := io.ReadAll(body)

		if err != nil { http.Error(w, "putRoom(): could not read body" , 500); return }
		err = json.Unmarshal(data, &requestData)
		if err != nil { http.Error(w, "putRoom(): could not unmarshal full", 500); return }

		ridInt, err := strconv.ParseInt(requestData.Data.Rid, 10, 32)
		if err != nil { http.Error(w, "putRoom(): could not convert parse rid", 500); return }
		uidInt, err := strconv.ParseInt(requestData.Data.Uid, 10, 32)
		if err != nil { http.Error(w, "putRoom(): could not convert parse uid", 500); return }

		// add user to room or send a message to room
		currentTimeStamp := time.Now().Unix()
		currentTimeStr := time.Unix(currentTimeStamp, 0).UTC().Format("2006-01-02 15:04:05-07")
		switch requestData.DataType {
		case "message":
			ms := s.MessageStore
			stat, err := ms.AddMessage(ctx, currentTimeStr, requestData.Data.Content, uint32(uidInt), uint32(ridInt))
			if err != nil { 
				fmt.Println(err)
				http.Error(w, "putRoom(): could not add message", 500)
				return 
			}
			fmt.Fprintf(w, "%v", stat)
		case "user":
			rs := s.RoomStore

			stat, err := rs.AddRoomUser(ctx, currentTimeStr, uint32(ridInt), uint32(uidInt))
			if err != nil { http.Error(w, "putRoom(): could not add user", 500); return }

			fmt.Fprintf(w, "%v", stat)
		default:
			http.Error(w, "putRoom(): invalid data type", 500)
			return
		}
	}
}
