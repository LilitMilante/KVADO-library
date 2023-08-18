-- +goose Up
-- +goose StatementBegin
CREATE TABLE authors
(
    id VARCHAR(36) PRIMARY KEY,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE authors;
-- +goose StatementEnd
