-- +goose Up
-- +goose StatementBegin
CREATE TABLE rfids (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    card_id VARCHAR(255) NOT NULL,
    user_id INT NOT NULL,
    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE rfids;
-- +goose StatementEnd
