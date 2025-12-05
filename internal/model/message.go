package model

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Message struct {
	id uint16
	time string
	content string
	userId uint16
	roomId uint16
}

type MessageStore interface {
	GetMessage(strat GetMessageStrategy) ([]Message, error)
	AddMessage(time, content string, userId, roomId uint16) (string, error)
	DeleteMessage(id uint16) (string, error)
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
	userId, roomId uint16
}

func (strat *GetMessageByRoomUser) MakeQuery() (string, []any) {
	return "select * from messages where user_id = $1 and room_id = $2", []any{strat.userId, strat.roomId}
}

type GetMessageByUser struct {
	UserId uint16
}

func (strat *GetMessageByUser) MakeQuery() (string, []any) {
	return "select * from messages where user_id = $1", []any{strat.UserId}
}

func (store *PgMessageStore) GetMessage(msgStrat GetMessageStrategy) ([]Message, error) {
	query, args := msgStrat.MakeQuery()
	rows, err := store.Db.Query(store.Ctx, query, args...)
	if err != nil { return nil, err }
	defer rows.Close()

	var msgs []Message

	for rows.Next() {
		var msg Message
		rows.Scan(&msg.id, &msg.time, &msg.content, &msg.userId, &msg.roomId)
		msgs = append(msgs, msg)
	}

	return msgs, nil
}

// other functions

func (store *PgMessageStore) AddMessage(time, content string, userId, roomId uint16) (string, error) {
	tag, err := store.Db.Exec(store.Ctx,
		"insert into message (time, content, user_id, room_id) " +
		"values ($1, $2, $3, $4)",
		time, content, userId, roomId,
	)
	if err != nil { return "", err }
	return tag.String(), nil
}

func (store *PgMessageStore) DeleteMessage(id uint16) (string, error) {
	tag, err := store.Db.Exec(store.Ctx, 
		"delete from messages " +
		"where id = $1",
		id,
	)
	if err != nil { return "", err }
	return tag.String(), nil
}
