package model

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Room struct {
	Id uint32
	Time time.Time
	Name string
}

type RoomStore interface {
	GetRoom(ctx context.Context, id uint32) (Room, error)
	GetAllRooms(ctx context.Context) ([]Room, error)
	AddRoom(ctx context.Context, time, name string) (string, error)
	AddRoomUser(ctx context.Context, time string, rid, uid uint32) (string, error)
	DeleteRoom(ctx context.Context, id uint32) (string, error)
}

type PgRoomStore struct {
	Db *pgxpool.Pool
}

func (store *PgRoomStore) GetAllRooms(ctx context.Context) ([]Room, error) {
	rows, err := store.Db.Query(ctx, "select * from rooms")
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

func (store *PgRoomStore) GetRoom(ctx context.Context, id uint32) (Room, error) {
	rows, err := store.Db.Query(ctx, "select * from rooms where id = $1", id)
	if err != nil { return Room{}, err }
	defer rows.Close()

	var rooms []Room
		
	for rows.Next() {
		var room Room

		rows.Scan(&room.Id, &room.Time, &room.Name)
		rooms = append(rooms, room)
	}
	
	if len(rooms) == 0 || rooms == nil {
		return Room{}, errors.New("no value from query");
	}

	return rooms[0], nil
}

func (store *PgRoomStore) AddRoom(ctx context.Context, time, name string) (string, error) {
	tag, err := store.Db.Exec(ctx, 
		"insert into rooms (time, name) " + 
		"values ($1, $2)",
		time, name,
	)
	if err != nil { return "", err }
	return tag.String(), nil
}

func (store *PgRoomStore) AddRoomUser(ctx context.Context, time string, rid, uid uint32) (string, error) {
	tag, err := store.Db.Exec(
		ctx,
		"insert into user_room_join (time, user_id, room_id) " +
		"values ($1, $2, $3)",
		time, uid, rid,
	)
	if err != nil { return "", err }
	return tag.String(), nil
}

func (store *PgRoomStore) DeleteRoom(ctx context.Context, id uint32) (string, error) {
	tag, err := store.Db.Exec(
		ctx,
		"delete from rooms " +
		"where id = $1", id,
	)
	if err != nil { return "", err }
	return tag.String(), nil
}
