-- +goose Up
-- +goose StatementBegin
CREATE TABLE authors_books
(
    author_id VARCHAR(36),
    book_id VARCHAR(36),
    PRIMARY KEY (author_id, book_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE authors_books;
-- +goose StatementEnd

