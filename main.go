package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/Aohk22/web-2-go-crud-msg/internal/model"
	"github.com/Aohk22/web-2-go-crud-msg/internal/srv"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func run(ctx context.Context, logger *log.Logger) error {
	// load envs
	err := godotenv.Load()
	if err != nil { return err }
	dbUrl, exist := os.LookupEnv("DATABASE_URL")
	if !exist { return errors.New("Env DATABASE_URL not set.") }

	// init db
	dbPool, err := pgxpool.New(ctx, dbUrl)
	if err != nil { return err }
	defer dbPool.Close()

	userStore := &model.PgUserStore{ Ctx: ctx, Db: dbPool }

	mux := srv.NewServer(logger, userStore)

	server := &http.Server{
		Addr: ":8080",
		Handler: mux,
	}

	go func() {
		log.Printf("Listening on %s\n", server.Addr)
		err := server.ListenAndServe()
		if err != nil { logger.Print("Could not init server") }
	}()
	var wg sync.WaitGroup
	wg.Go(func() {
		<-ctx.Done()
		shutdownCtx, end := context.WithTimeout(context.Background(), 10 * time.Second)
		defer end()
		if err := server.Shutdown(shutdownCtx); err != nil {
			logger.Printf("error shutting down server %s\n", err)
		}
	})
	wg.Wait()
	return nil
}

func main() {
	ctx := context.Background()
	logger := log.New(os.Stdout, "MyLoggER: ", log.Ldate)

	run(ctx, logger)
}

