CREATE DATABASE chatapi;


CREATE TABLE users (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    username VARCHAR NOT NULL UNIQUE,
    created_at TIMESTAMP
);

CREATE TABLE chats (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    name VARCHAR NOT NULL,
    created_at TIMESTAMP
);

CREATE TABLE messages (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    chat_id INTEGER REFERENCES chats(id),
    author INTEGER REFERENCES users(id),
    text VARCHAR,
    created_at TIMESTAMP
);

CREATE TABLE chat_users (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    chat_id INTEGER REFERENCES chats(id),
    user_id INTEGER REFERENCES users(id)
);