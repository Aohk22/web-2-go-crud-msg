package model

import (
	"context"
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
	GetMessage(ctx context.Context, strat GetMessageStrategy) ([]Message, error)
	AddMessage(ctx context.Context, time, content string, userId, roomId uint32) (string, error)
	DeleteMessage(ctx context.Context, id uint32) (string, error)
}

type PgMessageStore struct {
	Db *pgxpool.Pool
}

// message get strategies

type GetMessageStrategy interface {
	MakeQuery() (string , []any)
}

// better off using inline functions probably.
type GetMessageByRoomUser struct {
	userId, roomId uint32
}

type GetMessageByUser struct {
	UserId uint32
}

type GetMessageByRoom struct {
	RoomId uint32
	Time time.Time
}

func (strat *GetMessageByRoomUser) MakeQuery() (string, []any) {
	return "select * from messages where user_id = $1 and room_id = $2", []any{strat.userId, strat.roomId}
}

func (strat *GetMessageByUser) MakeQuery() (string, []any) {
	return "select * from messages where user_id = $1", []any{strat.UserId}
}

func (strat *GetMessageByRoom) MakeQuery() (string, []any) {
	query :=
		`select *
		from messages
		where room_id = $1
		and time < $2
		order by time desc
		limit 5`
		
	return query, []any{strat.RoomId, strat.Time}
}

func (store *PgMessageStore) GetMessage(ctx context.Context, msgStrat GetMessageStrategy) ([]Message, error) {
	query, args := msgStrat.MakeQuery()
	rows, err := store.Db.Query(ctx, query, args...)
	if err != nil { return nil, err }
	defer rows.Close()

	var msgs []Message

	for rows.Next() {
		var msg Message
		rows.Scan(&msg.Id, &msg.Time, &msg.Content, &msg.UserId, &msg.RoomId)
		msgs = append(msgs, msg)
	}

	return msgs, nil
}

// other functions

func (store *PgMessageStore) AddMessage(ctx context.Context, time, content string, userId, roomId uint32) (string, error) {
	tag, err := store.Db.Exec(
		ctx,
		"insert into messages (time, content, user_id, room_id) " +
		"values ($1, $2, $3, $4)",
		time, content, userId, roomId,
	)
	if err != nil { return "", err }
	return tag.String(), nil
}

func (store *PgMessageStore) DeleteMessage(ctx context.Context, id uint32) (string, error) {
	tag, err := store.Db.Exec(
		ctx, 
		"delete from messages " +
		"where id = $1",
		id,
	)
	if err != nil { return "", err }
	return tag.String(), nil
}
