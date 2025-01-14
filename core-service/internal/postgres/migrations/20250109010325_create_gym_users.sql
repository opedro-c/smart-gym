-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    enabled BOOLEAN DEFAULT TRUE NOT NULL,
    username VARCHAR(25) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    password TEXT NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
