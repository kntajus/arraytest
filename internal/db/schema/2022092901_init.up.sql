CREATE TYPE fruit AS ENUM ('apple', 'banana', 'kiwi');

CREATE TABLE choices (
    choice_id serial PRIMARY KEY,
    fruits fruit[] NOT NULL
);
