CREATE TABLE IF NOT EXISTS cards (
    card_id serial PRIMARY KEY,
    card_name VARCHAR (200) NOT NULL UNIQUE,
    card_number VARCHAR (200) NOT NULL,
    owner VARCHAR (200) NOT NULL,
    cvv VARCHAR (200) NOT NULL,
    card_date VARCHAR(200) NOT NULL,
    created_time TIMESTAMP NOT NULL
    );