-- Migration: create
-- Created at: 2018-09-16 17:21:28
-- ====  UP  ====

BEGIN;
	PRAGMA foreign_keys = ON;

	CREATE TABLE rooms (
		room_id integer,
		owner_id integer NOT NULL,
		name text UNIQUE NOT NULL,
		media_type text,
		media_source text,
		
		created_at text NOT NULL DEFAULT CURRENT_TIMESTAMP,
		modified_at text NOT NULL DEFAULT CURRENT_TIMESTAMP,

		PRIMARY KEY (room_id),
		FOREIGN KEY (owner_id) REFERENCES users(user_id)
	);

	CREATE TABLE users (
		user_id integer,
		name text UNIQUE NOT NULL,
		email text UNIQUE,
		token blob,
		level integer,

		created_at text NOT NULL DEFAULT CURRENT_TIMESTAMP,
		
		PRIMARY KEY (user_id)
	);

	INSERT INTO users (user_id, name, level) VALUES 
		(0, 'anonymous', null),
		(1, 'root', 10);
	
	-- create+reserve lobby
	INSERT INTO rooms (
		room_id, owner_id, name, 
		media_type, media_source
	) VALUES (
		0, 1, 'lobby', 
		'text', null
	);

COMMIT;

-- ==== DOWN ====

BEGIN;
	DROP TABLE rooms;
	DROP TABLE users;
COMMIT;
