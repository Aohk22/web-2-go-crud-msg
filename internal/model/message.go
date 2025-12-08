package model

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Message struct {
	Id uint32
	Time time.Time
	Content string
	UserId uint32
	RoomId uint32
}

type MessageStore interface {
	GetMessage(strat GetMessageStrategy) ([]Message, error)
	AddMessage(time, content string, userId, roomId uint32) (string, error)
	DeleteMessage(id uint32) (string, error)
}

type PgMessageStore struct {
	Ctx context.Context
	Db *pgxpool.Pool
}

// message get strategies

type GetMessageStrategy interface {
	MakeQuery() (string , []any)
}

type GetMessageByRoomUser struct {
	userId, roomId uint32
}

type GetMessageByUser struct {
	UserId uint32
}

type GetMessageByRoom struct {
	RoomId uint32
}

func (strat *GetMessageByRoomUser) MakeQuery() (string, []any) {
	return "select * from messages where user_id = $1 and room_id = $2", []any{strat.userId, strat.roomId}
}

func (strat *GetMessageByUser) MakeQuery() (string, []any) {
	return "select * from messages where user_id = $1", []any{strat.UserId}
}

func (strat *GetMessageByRoom) MakeQuery() (string, []any) {
	return "select * from messages where room_id = $1", []any{strat.RoomId}
}

func (store *PgMessageStore) GetMessage(msgStrat GetMessageStrategy) ([]Message, error) {
	query, args := msgStrat.MakeQuery()
	rows, err := store.Db.Query(store.Ctx, query, args...)
	if err != nil { return nil, err }
	defer rows.Close()

	var msgs []Message

	for rows.Next() {
		var msg Message
		rows.Scan(&msg.Id, &msg.Time, &msg.Content, &msg.UserId, &msg.RoomId)
		log.Println(msg)
		msgs = append(msgs, msg)
	}

	return msgs, nil
}

// other functions

func (store *PgMessageStore) AddMessage(time, content string, userId, roomId uint32) (string, error) {
	tag, err := store.Db.Exec(store.Ctx,
		"insert into message (time, content, user_id, room_id) " +
		"values ($1, $2, $3, $4)",
		time, content, userId, roomId,
	)
	if err != nil { return "", err }
	return tag.String(), nil
}

func (store *PgMessageStore) DeleteMessage(id uint32) (string, error) {
	tag, err := store.Db.Exec(store.Ctx, 
		"delete from messages " +
		"where id = $1",
		id,
	)
	if err != nil { return "", err }
	return tag.String(), nil
}
