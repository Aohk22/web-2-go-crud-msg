create database msgapp;

create table users (
	id serial primary key,
	name varchar(255),
	passhash varchar(255) -- bycrypt hashed
);

create table rooms (
	id serial primary key,
	time timestamptz,
	name varchar(255)
);

create table user_room_join (
	time timestamptz,

	user_id integer references users (id),
	room_id integer references rooms (id)
);

create table messages (
	id serial primary key,
	time timestamptz,
	content text,

	user_id integer references users (id),
	room_id integer references rooms (id)
);
