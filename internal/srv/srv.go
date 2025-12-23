package srv

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Aohk22/web-2-go-crud-msg/internal/model"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var jwtKey = []byte(os.Getenv("JWT_KEY"))

type Stores struct {
	UserStore model.UserStore
	RoomStore model.RoomStore
	MessageStore model.MessageStore
}

func NewStores(db *pgxpool.Pool) (*Stores) {
	return &Stores{
		UserStore: &model.PgUserStore { Db: db },
		RoomStore: &model.PgRoomStore { Db: db },
		MessageStore: &model.PgMessageStore { Db: db },
	}
}

func NewServer(ctx context.Context, logger *log.Logger, db *pgxpool.Pool) (http.Handler, error) {
	var handler http.Handler
	var mux *http.ServeMux = http.NewServeMux()

	_, exist := os.LookupEnv("JWT_KEY")
	if !exist { return nil, errors.New("Env JWT_KEY not set.") }

	stores := NewStores(db)
	hub := NewHub()
	go hub.run()

	addRoutes(ctx, mux, stores, hub)

	handler = authMiddleware(mux)
	handler = corsMiddleware(handler)
	handler = loggerMiddleware(logger, handler)

	return handler, nil
}

func createToken(username string, uid string) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": username, // name
		"uid": uid, // uid
		"iss": "msgapp", // issuer
		"exp": time.Now().Add(time.Hour).Unix(), // expiry
		"iat": time.Now().Unix(), // issued at
	})

	tokenString, err := claims.SignedString(jwtKey)
	if err != nil { return  "", err }

	return tokenString, nil
}

func verifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		return jwtKey, nil
	})
	if err != nil { return nil, err }
	if !token.Valid { return nil, fmt.Errorf("invalid token") }
	return token, nil
}
