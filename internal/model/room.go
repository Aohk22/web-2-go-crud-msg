package model

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Room struct {
	Id uint32
	Time time.Time
	Name string
}

type RoomStore interface {
	GetRoom(id uint32) ([]Room, error)
	GetAllRooms() ([]Room, error)
	AddRoom(time, name string) (string, error)
	DeleteRoom(id uint32) (string, error)
}

type PgRoomStore struct {
	Ctx context.Context
	Db *pgxpool.Pool
}

func (store *PgRoomStore) GetAllRooms() ([]Room, error) {
	rows, err := store.Db.Query(store.Ctx, "select * from rooms")
	if err != nil { return nil, err }
	defer rows.Close()

	var rooms []Room

	for rows.Next() {
		var room Room

		rows.Scan(&room.Id, &room.Time, &room.Name)
		rooms = append(rooms, room)
	}

	return rooms, nil
}

func (store *PgRoomStore) GetRoom(id uint32) ([]Room, error) {
	rows, err := store.Db.Query(store.Ctx, "select * from rooms where id = $1", id)
	if err != nil { return nil, err }
	defer rows.Close()

	var rooms []Room
		
	for rows.Next() {
		var room Room

		rows.Scan(&room.Id, &room.Time, &room.Name)
		rooms = append(rooms, room)
	}

	return rooms, nil
}

func (store *PgRoomStore) AddRoom(time, name string) (string, error) {
	tag, err := store.Db.Exec(store.Ctx, 
		"insert into rooms (time, name) " + 
		"values ($1, $2)",
		time, name,
	)
	if err != nil { return "", err }
	return tag.String(), nil
}

func (store *PgRoomStore) DeleteRoom(id uint32) (string, error) {
	tag, err := store.Db.Exec(store.Ctx,
		"delete from rooms " +
		"where id = $1", id,
	)
	if err != nil { return "", err }
	return tag.String(), nil
}
