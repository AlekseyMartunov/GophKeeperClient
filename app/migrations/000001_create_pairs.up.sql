CREATE TABLE IF NOT EXISTS pairs (
    pair_id serial PRIMARY KEY,
    pair_name VARCHAR (200) NOT NULL UNIQUE,
    password VARCHAR (200) NOT NULL,
    login VARCHAR (200) NOT NULL,
    created_time TIMESTAMP NOT NULL
);