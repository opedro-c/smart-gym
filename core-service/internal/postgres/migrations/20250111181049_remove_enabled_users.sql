-- +goose Up
-- +goose StatementBegin
ALTER TABLE users
    DROP COLUMN enabled;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE users
    ADD enabled BOOLEAN DEFAULT TRUE NOT NULL;
-- +goose StatementEnd
