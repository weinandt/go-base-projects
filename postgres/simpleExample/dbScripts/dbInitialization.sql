CREATE DATABASE test_db;

-- Psql to switch to the new database.
\c test_db

CREATE EXTENSION "uuid-ossp";

CREATE TABLE IF NOT EXISTS users (
    id uuid NOT NULL DEFAULT uuid_generate_v4(),
    name text NOT NULL,
    CONSTRAINT user_pk PRIMARY KEY (id)
);

INSERT INTO users (name) VALUES ('nick'), ('anotherNick');
