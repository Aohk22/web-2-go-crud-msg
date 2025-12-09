package srv

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/Aohk22/web-2-go-crud-msg/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type TestDefinition struct {
	Method string
	Path string
	Want any
	Getter func(io.ReadCloser) (any, error)
}

const testDbUrl = "postgresql://tkl:123@localhost:5432/msgapp"
const fatFmt = "Want: %v Got: %v\n"
var serveMux *http.ServeMux

var testTable = []TestDefinition { 
	 {
		Method: "GET", Path: "/users", Want: 5, Getter: func(body io.ReadCloser) (any, error) {
			var users []model.User
			err := getData(body, &users)
			return len(users), err
		},
	},
	{
		Method: "GET", Path: "/user/1", Want: "Alice",
		Getter: func(body io.ReadCloser) (any, error) {
			var user model.User
			err := getData(body, &user)
			return user.Name, err
		},
	},
	{
		Method: "GET", Path: "/user/1/messages", Want: "2025-11-15 17:36:12 +0700 +07", 
		Getter: func(body io.ReadCloser) (any, error) {
			var messages []model.Message
			err := getData(body, &messages)
			return (messages[0].Time).String(), err
		},
	},
	{
		Method: "GET", Path: "/rooms", Want: 5, 
		Getter: func(body io.ReadCloser) (any, error) {
			var rooms []model.Room
			err := getData(body, &rooms)
			return len(rooms), err
		},
	},
	{
		Method: "GET", Path: "/room/1", Want: "general",
		Getter: func(body io.ReadCloser) (any, error) {
			var room model.Room
			err := getData(body, &room)
			return room.Name, err
		},
	},
	{
		Method: "GET", Path: "/room/1/messages", Want: "wrong want",
		Getter: func(body io.ReadCloser) (any, error) {
			var messages []model.Message
			err := getData(body, &messages)
			return (messages[0].Time).String(), err
		},
	},
	{
		Method: "GET", Path: "/room/1/users", Want: "wrong want",
		Getter: func(body io.ReadCloser) (any, error) {
			var users []model.User
			err := getData(body, &users)
			return users, err
		},
	},
}


func TestMain(m *testing.M) {
	ctx := context.Background()
	dbPool, _ := pgxpool.New(context.Background(), testDbUrl)

	stores := NewStores(dbPool)
	serveMux = http.NewServeMux()

	addRoutes(ctx, serveMux, stores)
	exitCode := m.Run()
	os.Exit(exitCode)
}


func TestAll(t *testing.T) {
	for _, tt := range testTable {
		t.Run("Test: " + tt.Method + tt.Path, func(t *testing.T) {
			req := httptest.NewRequest(tt.Method, tt.Path, nil)
			rec := httptest.NewRecorder()

			serveMux.ServeHTTP(rec, req)
			res := rec.Result()
			defer res.Body.Close()

			got, err := tt.Getter(res.Body)
			if err != nil { t.Error(err) }

			if tt.Want != got { t.Fatalf(fatFmt, tt.Want, got) }
		})
	}
}


func getData(resBody io.ReadCloser, v any) (error) {
	data, err := io.ReadAll(resBody)
	if err != nil { return err }
	err = json.Unmarshal(data, v)
	if err != nil { return err }
	return nil
}
