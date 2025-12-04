package main

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/Aohk22/2-go-crud-msg/srv"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func run(ctx context.Context, pool *pgxpool.Pool, w io.Writer, args []string) error {

}

func main() {
	err := godotenv.Load()
	if err != nil { fmt.Fprintf(os.Stderr, "%s\n", err) }

	dbUrl, exist := os.LookupEnv("DATABASE_URL")
	if !exist { fmt.Fprintf(os.Stderr, "env DATABASE_URL not set") }

	ctx := context.Background()

	dbPool, err := pgxpool.New(ctx, dbUrl)
	if err != nil { fmt.Fprintf(os.Stderr, "%s\n", err) }
	defer dbPool.Close()

	run(ctx, dbPool, os.Stdout, os.Args)
}

