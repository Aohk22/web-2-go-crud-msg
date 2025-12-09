package srv

import (
	"context"
	"log"
	"net/http"

	"github.com/Aohk22/web-2-go-crud-msg/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

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

func NewServer(ctx context.Context, logger *log.Logger, db *pgxpool.Pool) http.Handler {
	var handler http.Handler
	var mux *http.ServeMux = http.NewServeMux()

	stores := NewStores(db)

	addRoutes(ctx, mux, stores)

	// TODO: wrap the loggers and auth middleware
	handler = loggerMiddleware(logger)(mux)

	return handler
}

