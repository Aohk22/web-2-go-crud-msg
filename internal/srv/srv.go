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

func NewStores(ctx context.Context, db *pgxpool.Pool) (*Stores) {
	return &Stores{
		UserStore: &model.PgUserStore { Ctx: ctx, Db: db },
		RoomStore: &model.PgRoomStore { Ctx: ctx, Db: db },
		MessageStore: &model.PgMessageStore { Ctx: ctx, Db: db },
	}
}

func NewServer(ctx context.Context, logger *log.Logger, db *pgxpool.Pool) http.Handler {
	var handler http.Handler
	var mux *http.ServeMux = http.NewServeMux()

	stores := NewStores(ctx, db)

	addRoutes(mux, stores)

	// TODO: wrap the loggers and auth middleware
	handler = loggerMiddleware(logger)(mux)

	return handler
}

