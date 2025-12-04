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
	GetRoom()
	AddRoom()
	RemoveRoom()
}

type PgRoomStore struct {
	ctx context.Context
	db *pgxpool.Pool
}


