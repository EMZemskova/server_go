-- +goose Up
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username TEXT UNIQUE NOT NULL,
    password TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS messages (
    id SERIAL PRIMARY KEY,
    chat INT NOT NULL,
    sender INT NOT NULL,
    text TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS chats (
    id SERIAL PRIMARY KEY,
    creator INT NOT NULL,
    guest INT NOT NULL,
    status TEXT NOT NULL
);

-- +goose Down
DROP TABLE IF EXISTS chats;
DROP TABLE IF EXISTS messages;
DROP TABLE IF EXISTS users;