-- +goose Up
CREATE TABLE chat (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    body TEXT,
    creation_date TIMESTAMP NOT NULL DEFAULT NOW(),
    update_date TIMESTAMP,
    deleted boolean NOT NULL DEFAULT false
);

CREATE TABLE messages (
    id SERIAL PRIMARY KEY,
    text TEXT,
    chat_id INT REFERENCES CHAT (ID),
    user_id INT NOT NULL,
    sent_date TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE links (
    chat_id INT REFERENCES CHAT (ID),
    user_id INT,
    invite_date TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted boolean NOT NULL DEFAULT false
);

-- +goose Down
DROP TABLE chat;
DROP TABLE messages;
DROP TABLE links;
