package db

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

type DatabaseHandler struct {
	ctx context.Context
	pool *pgxpool.Pool
	logger *log.Logger
}

func (db *DatabaseHandler) Query(query string) (*pgx.Rows) {
	rows, err := db.pool.Query(db.ctx, query)
	if err != nil { db.logger.Fatal(err) }
	return &rows
}

func (db *DatabaseHandler) Exec(query string) (string) {
	tag, err := db.pool.Exec(db.ctx, query)
	if err != nil { db.logger.Fatal(err) }
	return tag.String()
}

func (db *DatabaseHandler) PrintAllUsers() {
	query := "select id, name, passhash from users limit 10;"

	rows := db.Query(query)
	defer (*rows).Close()

	for (*rows).Next() {
		values, err := (*rows).Values()
		if err != nil { log.Fatal("printAllUsers row.Values() error. ", err) }
		fmt.Println(values)
	}
}

func (db *DatabaseHandler) AddUser(name string, password []byte) {
	// process
	hash, err := bcrypt.GenerateFromPassword(password, 10)
	if err != nil { log.Fatal("addUser bcrypt error. ", err) }

	// create query
	query := fmt.Sprintf("insert into users(name, passhash) values(%s, %s);", name, hash)

	tagStr := db.Exec(query)
	db.logger.Println(tagStr)
}

func (db *DatabaseHandler) RemoveUser(pool *pgxpool.Pool, id uint16) {
	query := fmt.Sprintf("delete from users where id = $1;", id)

	tagStr := db.Exec(query)
	db.logger.Println(tagStr)
}

