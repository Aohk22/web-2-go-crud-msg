package model

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

// defs 

type User struct {
	id uint16
	Name string
	Password []byte
}

type UserStore interface {
	GetAllUsers() ([]User, error)
	GetUser(id uint16) (User, error)
	AddUser(name string, password []byte) (string, error)
	RemoveUser(id uint16) (string, error)
}

// implementation

type PgUserStore struct {
	Ctx context.Context
	Db *pgxpool.Pool
}

func (s *PgUserStore) GetAllUsers() ([]User, error) {
	rows, err := s.Db.Query(s.Ctx, "select * from users;")
	if err != nil { return nil, err }
	defer rows.Close()

	var users []User

	for rows.Next() {
		var u User

		rows.Scan(&u.id, &u.Name, &u.Password)
		users = append(users, u)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (s *PgUserStore) GetUser(id uint16) (User, error) {
	rows, err := s.Db.Query(s.Ctx, "select * from users where id = $1", id)
	if err != nil { return User{}, err }

	var user User
	
	for rows.Next() {
		rows.Scan(&user.id, &user.Name, &user.Password)
		break // since id is unique, assume 1
	}
	return user, nil
}

func (m *PgUserStore) AddUser(name string, password []byte) (string, error) {
	passhash, err := bcrypt.GenerateFromPassword(password, 10)
	if err != nil { return "", err }
	tag, err := m.Db.Exec(m.Ctx, "insert into users (name, passhash) values ($1, $2);", name, passhash)
	if err != nil { return "", err }
	return tag.String(), nil
}

func (m *PgUserStore) RemoveUser(id uint16) (string, error) {
	tag, err := m.Db.Exec(m.Ctx, "delete from users where id = $1;", id)
	if err != nil { return "", err }
	return tag.String(), nil
}
