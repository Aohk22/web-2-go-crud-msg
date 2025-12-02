package main

import (
	"os"
	"log"
	"context"

	"github.com/joho/godotenv"
	"github.com/jackc/pgx/v5/pgxpool"
)

const STDIN = 0

var pool *pgxpool.Pool

func main() {
	var err error
	var dbUrl string
	var pool *pgxpool.Pool

	err = godotenv.Load()
	if err != nil { log.Fatal(err) }

	dbUrl = os.Getenv("DATABASE_URL")
	if len(dbUrl) == 0 { log.Fatal("env DATABASE_URL not found.") }

	pool, err = pgxpool.New(context.Background(), dbUrl)
	if err != nil { log.Fatal(err) }
	defer pool.Close()

}

