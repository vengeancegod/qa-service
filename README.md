QA Service API
Микросервис для управления вопросами и ответами с REST API, построенный на Go с чистой архитектурой.
```
Архитектура в проекте представлена следующим образом:
qa-service/
├── cmd/app/                 
├── internal/
│   ├── api/handlers/       # HTTP обработчики
│   ├── app/               # (Service Provider)
│   ├── config/            # Конфигурация (DB, HTTP)
│   ├── infrastructure/    # Инфраструктура (DB connection)
│   ├── model/             # Модели данных
│   ├── repository/        # Слой репозиториев
│   └── service/           # Слой бизнес-логики
├── migrations/
```

Запуск приложения:
```bash
git clone https://github.com/vengeancegod/qa-service.git
cd qa-service

docker compose up -d --build
```
Создание миграций:
```bash
make migrate-up
```
Запуск тестов:
```bash
make test
```
Работа с API:
GET /questions/ — список всех вопросов 
```bash
curl -X GET http://localhost:8081/questions/
```

POST /questions/ — создать новый вопрос
```bash
curl -X POST http://localhost:8081/questions/ \
  -H "Content-Type: application/json" \
  -d '{
    "text": "Текст?"
  }'
```

GET /questions/{id} — получить вопрос и все ответы на него
```bash
curl -X GET http://localhost:8081/questions/1/
```
DELETE /questions/{id} — удалить вопрос (вместе с ответами)
```bash
curl -X DELETE http://localhost:8081/questions/3/
```

POST /questions/{id}/answers/ — добавить ответ к вопросу
```bash
curl -X POST http://localhost:8081/questions/1/answers/ \
  -H "Content-Type: application/json" \
  -d '{
    "text": "Ответ"
  }'
```

GET /answers/{id} — получить конкретный ответ
```bash
curl -X GET http://localhost:8081/answers/1/
```

DELETE /answers/{id} — удалить ответ
```bash
curl -X DELETE http://localhost:8081/answers/3/
```
