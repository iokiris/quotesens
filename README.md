# Quotesens

Quotesens — это простой сервис для работы с цитатами. Поддерживает добавление, удаление, получение случайной или всех цитат.


- #### Архитектура сделана таким образом, что позволяет удобно масштабировать проект.
- #### В проекте учтены принципы программирования на Go
- #### Имеются кастомные технологические решения, например централизованный перехватчик ошибок.

## Старт проекта
### Требования:
- Docker
- Docker Compose

### Запуск:
- В корневой директории проекта:
    ```bash
    docker-compose up --build
    ```

- #### Для удобства создан index.html, чтобы не отправлять запросы вручную
- #### Сервер работает по адресу http://localhost:8080, порт меняется в .env

### Тестирование
- Для удобного тестирования добавлен файл с нужными запросами, запуск:
- ```bash
  bash curl.sh
  ```
  
#### Вывод с исходных запросов:
```bash
=== POST http://localhost:8080/quotes ===
Status: 201
Body:



=== GET http://localhost:8080/quotes ===
Status: 200
Body:
[{"ID":1,"Author":"Confucius","Quote":"Life is simple, but we insist on making it complicated."}]


=== GET http://localhost:8080/quotes/random ===
Status: 200
Body:
{"ID":1,"Author":"Confucius","Quote":"Life is simple, but we insist on making it complicated."}


=== GET http://localhost:8080/quotes?author=Confucius ===
Status: 200
Body:
[{"ID":1,"Author":"Confucius","Quote":"Life is simple, but we insist on making it complicated."}]


=== DELETE http://localhost:8080/quotes/1 ===
Status: 204
Body:



=== GET http://localhost:8080/quotes ===
Status: 200
Body:
null

```
## API

Метод | Endpoint | Описание
------|----------|---------
GET   | /quotes  | Получить все цитаты
GET   | /quotes/random | Получить случайную цитату
POST  | /quotes  | Добавить новую цитату
DELETE| /quotes/{id} | Удалить цитату по ID


## Конфигурация

- #### Порт по умолчанию: 8080

- #### Изменяется в docker-compose.yml:

    ```yml
    ports:
    - "8080:8080"
    ```

### Пример запроса
- ```bash
  curl -X POST http://localhost:8080/quotes \\
  -H "Content-Type: application/json" \\
  -d '{"author":"Лев Толстой", "text":"Все думают изменить мир, но никто не думает изменить себя."}'
  ```

