package model

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

// defs 

type User struct {
	id uint16
	name string
	passhash string
}

type UserStore interface {
	GetAllUsers() ([]User, error)
	AddUser(string, []byte) (string, error)
	RemoveUser(uint16) (string, error)
}

// implementation

type PgUserStore struct {
	Ctx context.Context
	Db *pgxpool.Pool
}

func (s *PgUserStore) GetAllUsers() ([]User, error) {
	rows, err := s.Db.Query(s.Ctx, "select * from users;")
	if err != nil { return nil, err }

	var users []User

	for rows.Next() {
		var u User

		rows.Scan(&u.id, &u.name, &u.passhash)
		users = append(users, u)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (m *PgUserStore) AddUser(name string, password []byte) (string, error) {
	passhash, err := bcrypt.GenerateFromPassword(password, 10)
	if err != nil { return "", err }
	tag, err := m.Db.Exec(m.Ctx, "insert into users (name, passhash) values ($1, $2);", name, passhash)
	if err != nil { return "", err }
	return tag.String(), nil
}

func (m *PgUserStore) RemoveUser(id uint16) (string, error) {
	tag, err := m.Db.Exec(m.Ctx, "delete from users where id = $1", id)
	if err != nil { return "", err }
	return tag.String(), nil
}
