QA Service API
Микросервис для управления вопросами и ответами с REST API, построенный на Go с чистой архитектурой.

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

