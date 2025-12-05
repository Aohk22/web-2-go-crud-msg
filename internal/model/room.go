package model

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Room struct {
	id uint16
	time string
	name string
}

type RoomStore interface {
	GetRoom(id uint16) ([]Room, error)
	AddRoom(time, name string) (string, error)
	DeleteRoom(id uint16) (string, error)
}

type PgRoomStore struct {
	Ctx context.Context
	Db *pgxpool.Pool
}

func (store *PgRoomStore) GetRoom(id uint16) ([]Room, error) {
	rows, err := store.Db.Query(store.Ctx, "select * from rooms where id = $1", id)
	if err != nil { return nil, err }
	defer rows.Close()

	var rooms []Room
		
	for rows.Next() {
		var room Room
		rows.Scan(&room.id, &room.time, &room.name)

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
func (store *PgRoomStore) DeleteRoom(id uint16) (string, error) {
	tag, err := store.Db.Exec(store.Ctx,
		"delete from rooms " +
		"where id = $1", id,
	)
	if err != nil { return "", err }
	return tag.String(), nil
}
