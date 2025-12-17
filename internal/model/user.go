package model

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id uint32
	Name string
	Password string
}

type UserStore interface {
	GetAllUsers(ctx context.Context) ([]User, error)

	GetUser(ctx context.Context, id uint32) (User, error)
	GetUsersByRoom(ctx context.Context, roomId uint32) ([]User, error)

	AddUser(ctx context.Context, name string, password string) (string, error)
	RemoveUser(ctx context.Context, id uint32) (string, error)

	CheckUser(ctx context.Context, name string, password string) (bool, error)
}

// implementation

type PgUserStore struct {
	Db *pgxpool.Pool
}

func (s *PgUserStore) GetAllUsers(ctx context.Context) ([]User, error) {
	rows, err := s.Db.Query(ctx, "select * from users;")
	if err != nil { return nil, err }
	defer rows.Close()

	var users []User

	for rows.Next() {
		var u User

		rows.Scan(&u.Id, &u.Name, &u.Password)
		users = append(users, u)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (s *PgUserStore) GetUser(ctx context.Context, id uint32) (User, error) {
	rows, err := s.Db.Query(ctx, "select * from users where id = $1", id)
	if err != nil { return User{}, err }
	defer rows.Close()

	var user User
	
	for rows.Next() {
		rows.Scan(&user.Id, &user.Name, &user.Password)
		break // since id is unique, assume 1
	}
	return user, nil
}

func (s *PgUserStore) GetUsersByRoom(ctx context.Context, roomId uint32) ([]User, error) {
	rows, err := s.Db.Query(
		ctx,
		`select u.id, u.name, u.passhash
		from rooms r
		join user_room_join ur on r.id = ur.room_id
		join users u on ur.user_id = u.id
		where r.id = $1`,
		roomId,
	)
	if err != nil { return nil, err }
	defer rows.Close()

	var users []User

	for rows.Next() {
		var user User
		err := rows.Scan(&user.Id, &user.Name, &user.Password)
		if err != nil { return nil, err }

		users = append(users, user)
	}

	return users, nil
}

func (m *PgUserStore) AddUser(ctx context.Context, name string, password string) (string, error) {
	passhash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil { return "", err }
	tag, err := m.Db.Exec(ctx, "insert into users (name, passhash) values ($1, $2);", name, passhash)
	if err != nil { return "", err }
	return tag.String(), nil
}

func (m *PgUserStore) RemoveUser(ctx context.Context, id uint32) (string, error) {
	tag, err := m.Db.Exec(ctx, "delete from users where id = $1;", id)
	if err != nil { return "", err }
	return tag.String(), nil
}

func (m *PgUserStore) CheckUser(ctx context.Context, name string, password string) (bool, error) {
	rows, err := m.Db.Query(ctx, "select passhash from users where name = $1", name)
	if err != nil { return false, err }
	defer rows.Close()

	var user User

	rows.Next()
	err = rows.Scan(&user.Password)
	if err != nil { return false, err }

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err == nil {
		return true, nil
	} else {
		return false, err 
	}
}
