-- CREATE DATABASE msgapp;

CREATE TABLE users (
	id serial primary key,
	name varchar(255),
	passhash varchar(255) -- bycrypt hashed
);

CREATE TABLE rooms (
	id serial primary key,
	time timestamptz,
	name varchar(255)
);

CREATE TABLE user_room_join (
	time timestamptz,

	user_id integer references users (id),
	room_id integer references rooms (id)
);

CREATE TABLE messages (
	id serial primary key,
	time timestamptz,
	content text,

	user_id integer references users (id),
	room_id integer references rooms (id)
);
