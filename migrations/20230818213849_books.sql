-- +goose Up
-- +goose StatementBegin
CREATE TABLE books
(
    id VARCHAR(36) PRIMARY KEY,
    title TEXT NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE books;
-- +goose StatementEnd
