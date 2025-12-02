package main

import (
	"os"
	"fmt"
	"log"
	"strconv"
	"context"

	"golang.org/x/term"

	"github.com/joho/godotenv"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/Aohk22/web-2-go-crud-msg/user"
)

const STDIN = 0

const (
	HELP = iota
	DISPLAY
	ADD
	REMOVE
	QUIT
)

func main() {
	err := godotenv.Load()
	if err != nil { log.Fatal(err) }

	dbUrl := os.Getenv("DATABASE_URL")

	pool, err := pgxpool.New(context.Background(), dbUrl)
	if err != nil { log.Fatal(err) }
	defer pool.Close()

	err = pool.Ping(context.Background())
	if err != nil { log.Fatal(err) }

	printHelp()

	run := true
	for run {
		var choice string

		fmt.Print("dbtest> ")
		fmt.Scanf("%s", &choice)

		opt, err := strconv.Atoi(choice)
		if err != nil { continue }

		switch opt {
		case HELP:
			printHelp()

		case DISPLAY:
			printAllUsers(pool)

		case ADD:
			var name string
			var pass []byte

			fmt.Print("User name to add: ")
			fmt.Scanf("%s", &name)
			fmt.Print("Users password: ")
			pass, err := term.ReadPassword(STDIN)
			if err != nil { log.Fatal(err) }
			
			addUser(pool, name, pass)

		case REMOVE:
			var id uint16

			fmt.Print("User id to remove: ")
			_, err := fmt.Scanf("%d", &id)
			if err != nil { log.Fatal(err) }

			removeUser(pool, id)

		case QUIT:
			run = false

		default:
			fmt.Println("invalid cmd")
		}
	}
}

func printHelp() {
	fmt.Println("--- help menu ---")
	fmt.Println("0 - help")
	fmt.Println("1 - display")
	fmt.Println("2 - add")
	fmt.Println("3 - remove")
	fmt.Println("4 - quit")
}

