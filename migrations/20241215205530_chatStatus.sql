-- +goose Up
CREATE TABLE example_table (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL
);

-- +goose Down
DROP TABLE example_table;
