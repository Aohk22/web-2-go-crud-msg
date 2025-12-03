package main

import (
	"os"
	"log"
	"context"

	"github.com/joho/godotenv"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	err := godotenv.Load()
	if err != nil { log.Fatal(err) }

	dbUrl, exist := os.LookupEnv("DATABASE_URL")
	if !exist { log.Fatal("env DATABASE_URL not set")}

	dbPool, err := pgxpool.New(context.Background(), dbUrl)
	if err != nil { log.Fatal(err) }

	defer dbPool.Close()
}

