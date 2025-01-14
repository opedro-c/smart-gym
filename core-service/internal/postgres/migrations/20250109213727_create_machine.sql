-- +goose Up
-- +goose StatementBegin
CREATE TABLE machines (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    name VARCHAR(25) NOT NULL,
    origin_id VARCHAR(255) NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE machines;
-- +goose StatementEnd
