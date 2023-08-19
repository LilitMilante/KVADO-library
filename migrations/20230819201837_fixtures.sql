-- +goose Up
-- +goose StatementBegin
-- Добавляем авторов
INSERT INTO authors (id, first_name, last_name)
VALUES ('16d6cb72-d119-4e2d-ab27-ff4940be0830', 'Александр', 'Пушкин'),
       ('44a54d7b-6289-4b12-b030-1ffd884763cb', 'Лев', 'Толстой'),
       ('37335b5a-fd7e-4fc3-9d67-7d981be01ba9', 'Федор', 'Достоевский');

-- Добавляем книги
INSERT INTO books (id, title)
VALUES ('6d33486d-acfe-433a-ae33-567b11c1e8e5', 'Евгений Онегин'),
       ('717525ad-e452-4cf1-99d2-fc3b2879055e', 'Война и мир'),
       ('53c82d2d-d435-4404-b151-a101067b489c', 'Преступление и наказание'),
       ('9e3a2d49-3eca-4951-adfd-6578fd22b02d', 'Анна Каренина'),
       ('f3abf142-715a-47a4-83da-4a681e24a278', 'Братья Карамазовы'),
       ('2fd328da-35e1-4394-b0e2-112082406b37', 'Вечный муж'),
       ('8e196342-e327-4ba6-b255-f5c3425bc2f4', 'Война и мир. Том 2'),
       ('ae66733b-7bdb-48bd-8adf-ce83a388c901', 'Анна Каренина. Часть 2'),
       ('4233b613-2f25-4ba7-94a8-516c88a44f1a', 'Преступление и наказание. Часть 2');

-- Добавляем авторам несколько книг
INSERT INTO authors_books (author_id, book_id)
VALUES ('16d6cb72-d119-4e2d-ab27-ff4940be0830', '6d33486d-acfe-433a-ae33-567b11c1e8e5'),
       ('44a54d7b-6289-4b12-b030-1ffd884763cb', '717525ad-e452-4cf1-99d2-fc3b2879055e'),
       ('37335b5a-fd7e-4fc3-9d67-7d981be01ba9', '53c82d2d-d435-4404-b151-a101067b489c'),
       ('44a54d7b-6289-4b12-b030-1ffd884763cb', '53c82d2d-d435-4404-b151-a101067b489c'),
       ('44a54d7b-6289-4b12-b030-1ffd884763cb', '9e3a2d49-3eca-4951-adfd-6578fd22b02d'),
       ('37335b5a-fd7e-4fc3-9d67-7d981be01ba9', 'f3abf142-715a-47a4-83da-4a681e24a278'),
       ('37335b5a-fd7e-4fc3-9d67-7d981be01ba9', '9e3a2d49-3eca-4951-adfd-6578fd22b02d'),
       ('44a54d7b-6289-4b12-b030-1ffd884763cb', 'f3abf142-715a-47a4-83da-4a681e24a278'),
       ('37335b5a-fd7e-4fc3-9d67-7d981be01ba9', '2fd328da-35e1-4394-b0e2-112082406b37');

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
TRUNCATE authors_books;
TRUNCATE books;
TRUNCATE authors;
-- +goose StatementEnd
