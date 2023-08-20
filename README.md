## Start the application with database:
`make up-docker`

## Install dependencies:
```bash
make deps
```


## How to send requests
### Get books by author ID:
```bash
make sand-1
```

### Get authors by book ID:
```bash
make send-2
```

# You can use next data examples

## Authors Table

| id                                   | first_name | last_name  |
| ------------------------------------ | ---------- | ---------- |
| 16d6cb72-d119-4e2d-ab27-ff4940be0830 | Александр  | Пушкин     |
| 44a54d7b-6289-4b12-b030-1ffd884763cb | Лев        | Толстой    |
| 37335b5a-fd7e-4fc3-9d67-7d981be01ba9 | Федор      | Достоевский|

## Books Table

| id                                   | title                     |
| ------------------------------------ | ------------------------- |
| 6d33486d-acfe-433a-ae33-567b11c1e8e5 | Евгений Онегин            |
| 717525ad-e452-4cf1-99d2-fc3b2879055e | Война и мир               |
| 53c82d2d-d435-4404-b151-a101067b489c | Преступление и наказание  |
| 9e3a2d49-3eca-4951-adfd-6578fd22b02d | Анна Каренина             |
| f3abf142-715a-47a4-83da-4a681e24a278 | Братья Карамазовы         |
| 2fd328da-35e1-4394-b0e2-112082406b37 | Вечный муж                |
| 8e196342-e327-4ba6-b255-f5c3425bc2f4 | Война и мир. Том 2        |
| ae66733b-7bdb-48bd-8adf-ce83a388c901 | Анна Каренина. Часть 2    |
| 4233b613-2f25-4ba7-94a8-516c88a44f1a | Преступление и наказание. Часть 2|
