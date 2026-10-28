CREATE TABLE items (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL
);

INSERT INTO items (name) VALUES ('item1'), ('item2'), ('item3');
