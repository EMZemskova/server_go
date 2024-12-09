CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name TEXT UNIQUE NOT NULL,
    password TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS messages (
    id SERIAL PRIMARY KEY,
    chat int64 NOT NULL,
    sender int64 NOT NULL,
    text TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS chats (
    id SERIAL PRIMARY KEY,
    creator int64 NOT NULL,
    guest int64 NOT NULL,
    status TEXT NOT NULL
);
