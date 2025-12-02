package user

import (
	"log"
	"fmt"
	"context"

	"golang.org/x/crypto/bcrypt"

	"github.com/jackc/pgx/v5/pgxpool"
)

func PrintAllUsers(pool *pgxpool.Pool) {
	query := "select id, name, passhash from users limit 10;"

	rows, err := pool.Query(context.Background(), query)
	if err != nil { log.Fatal("getAllUsers Query error. ", err) }
	defer rows.Close()

	for rows.Next() {
		values, err := rows.Values()
		if err != nil { log.Fatal("printAllUsers row.Values() error. ", err) }
		fmt.Println(values)
	}
}

func AddUser(pool *pgxpool.Pool, name string, password []byte) {
	query := "insert into users(name, passhash) values($1, $2);"

	hash, err := bcrypt.GenerateFromPassword(password, 10)
	if err != nil { log.Fatal("addUser bcrypt error. ", err) }


	_, err = pool.Exec(context.Background(), query, name, hash)
	if err != nil { log.Fatal("addUser Exec error. ", err) }
}

func RemoveUser(pool *pgxpool.Pool, id uint16) {
	query := "delete from users where id = $1;"

	_, err := pool.Exec(context.Background(), query, id)
	if err != nil { log.Fatal("removeUser Exec error. ", err) }
}

