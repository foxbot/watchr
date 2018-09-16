-- name: create-user
INSERT INTO users (name, token) VALUES (?, ?);

-- name: validate-token
SELECT ? = (
    SELECT (token)
    FROM users 
    WHERE name = ?
);

-- name: find-user-by-token
SELECT (user_id, name, email, level, created_at)
FROM users
WHERE token = ?;

-- name: create-room
INSERT INTO rooms (
    owner_id, name
) VALUES (
    ?, ?
);

-- name: find-room
SELECT (owner_id, name, media_type, media_source, created_at, modified_at)
FROM rooms
WHERE name = ?;

-- name: delete-room
DELETE FROM rooms
WHERE name = ?;