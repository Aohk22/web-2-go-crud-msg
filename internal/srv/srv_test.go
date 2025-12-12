package srv

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/Aohk22/web-2-go-crud-msg/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

type TestDefinition struct {
	Method string
	Path string
	Want any
	Getter func(io.ReadCloser) (any, error)
}

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
		Method: "GET", Path: "/room/1/messages", Want: []string{"3","2025-11-15 17:36:12 +0700 +07"},
		Getter: func(body io.ReadCloser) (any, error) {
			var messages []model.Message
			err := getData(body, &messages)
			return []string{fmt.Sprint(len(messages)), (messages[0].Time).String()}, err
		},
	},
	{
		Method: "GET", Path: "/room/1/users", 
		Want: []model.User{
			{Id: 1, Name: "Alice", Password: "$2b$12$Kixh5eX5z3fZ8jQ9vN7m5O8pL9kJ7hG5fD3sA1bC9xE7vT5rW2qU6"},
			{Id: 2, Name: "Bob", Password: "$2b$12$LmN7pQ8rT2vX5yZ9aB3cD5fH7jK9mP1qS3tU5wY8zA2cE4gI6kM8o"},
			{Id: 3, Name: "Charlie", Password: "$2b$12$XyZ9aB3cD5fH7jK9mP1qS3tU5wY8zA2cE4gI6kM8oP0qR2sT4uV6wX8"},
		},
		Getter: func(body io.ReadCloser) (any, error) {
			var users []model.User
			err := getData(body, &users)
			return users, err
		},
	},
}


func TestMain(m *testing.M) {
	ctx := context.Background()
	err := godotenv.Load()
	if err != nil { 
		fmt.Println(err) 
		os.Exit(1)
	}
	dbUrl, exist := os.LookupEnv("DATABASE_URL")
	if !exist {
		fmt.Println("Env DATABASE_URL not set.") 
		os.Exit(1)
	}
	dbPool, err := pgxpool.New(context.Background(), dbUrl)
	if err != nil { 
		fmt.Println(err)
		os.Exit(1)
	}

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

			switch typeof(got) {
			case "[]string":
				gl := got.([]string)
				for i, w := range gl {
					if gl[i] != w {
						t.Fatalf("Want: %v Got: %v\n", w, gl[i])
					}
				}
			case "[]model.User":
				gl := got.([]model.User)
				for i, w := range gl {
					if gl[i] != w {
						t.Fatalf("Want: %v Got: %v\n", w, gl[i])
					}
				}
			default:
				if tt.Want != got { t.Fatalf("Want: %v Got: %v\n", tt.Want, got) }
			}
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

func typeof(v any) string {
	return fmt.Sprintf("%T", v)
}
