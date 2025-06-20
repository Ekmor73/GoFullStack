# JobBoard

[![Go](https://img.shields.io/badge/Go-1.22-blue)](https://golang.org/) [![PostgreSQL](https://img.shields.io/badge/PostgreSQL-16-blue)](https://www.postgresql.org/) [![Docker](https://img.shields.io/badge/Docker-20.10-blue)](https://www.docker.com/) [![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

**JobBoard** — это веб-приложение для публикации и просмотра вакансий, разработанное на Go с использованием современного стека технологий, включая Fiber, Templ и HTMX. Проект создан как pet-проект для изучения веб-разработки на Go и демонстрации навыков backend-разработки. Приложение позволяет пользователям просматривать вакансии с пагинацией, создавать новые вакансии через форму и входить в систему с помощью сессионной аутентификации.

## Функциональность

- **Просмотр вакансий**: Отображение списка вакансий с пагинацией (2 вакансии на страницу).
- **Создание вакансий**: Форма для добавления вакансий с валидацией полей (роль, компания, зарплата, тип, локация, email).
- **Аутентификация**: Простая сессионная аутентификация с возможностью входа и выхода.
- **Динамический UI**: Использование HTMX для асинхронного обновления интерфейса без перезагрузки страницы.
- **Логирование**: Структурированное логирование через zerolog для отслеживания событий и ошибок.
- **Миграции базы данных**: SQL-скрипты для инициализации и обновления схемы базы данных.

## Технологии

- **Backend**:
  - Go (1.22)
  - Fiber (v2) — высокопроизводительный веб-фреймворк
  - pgx (v5) — драйвер для PostgreSQL
  - zerolog — быстрая библиотека для логирования
  - gobuffalo/validate — валидация данных
- **Frontend**:
  - Templ — генерация HTML-шаблонов на Go
  - HTMX (2.0) — для динамического взаимодействия
  - CSS — минимальная стилизация
- **База данных**: PostgreSQL (16)
- **Инфраструктура**:
  - Docker и Docker Compose — для контейнеризации
  - Air — для hot-reload во время разработки
  - Git — контроль версий

## Установка и запуск

### Требования

- Go (1.22 или выше)
- PostgreSQL (16 или выше)
- Docker и Docker Compose (опционально)
- Git

### Локальный запуск

1. **Клонируйте репозиторий**:
   ```bash
   git clone https://github.com/Ekmor73/GoFullStack.git
   cd GoFullStack

Настройте окружение:
Скопируйте .env.example в .env и укажите параметры:
env

DATABASE_URL=postgres://user:password@localhost:5432/jobboard
LOG_LEVEL=0
LOG_FORMAT=json

Примените миграции базы данных:
bash

psql -U postgres -d jobboard -f migrations/2025-05-21-1-init.sql
psql -U postgres -d jobboard -f migrations/2025-05-21-1-add-createdat.sql

Установите зависимости:
bash

go mod tidy

Запустите приложение:
Обычный запуск:
bash

go run cmd/main.go

С hot-reload через Air:
bash

air

Приложение будет доступно на http://localhost:3000.

Запуск через Docker
Убедитесь, что Docker и Docker Compose установлены.

Запустите приложение:
bash

docker-compose up --build

Приложение будет доступно на http://localhost:3000.

Пример использования
Откройте http://localhost:3000 в браузере, чтобы просмотреть список вакансий.

Перейдите на /login и войдите с тестовыми данными:
Email: a@a.ru

Пароль: 1

Создайте новую вакансию через форму на /vacancy.

Используйте кнопки пагинации для навигации по списку вакансий.

Планы улучшений
Реализовать регистрацию пользователей с хешированием паролей (bcrypt).

Добавить защиту от CSRF-атак для форм.

Внедрить юнит-тесты для ключевых компонентов (например, репозитория вакансий).

Задокументировать API с помощью Swagger.

Настроить CI/CD через GitHub Actions для автоматизации сборки и тестирования.

Развернуть демо-версию приложения на Render или Fly.io.

Добавить поддержку миграций через библиотеку golang-migrate.

Проблемы и ограничения
Текущая аутентификация использует жёстко закодированные данные (планируется переход на базу данных).

Отсутствует документация API (будет добавлена через Swagger).

Нет автоматических тестов (будут внедрены в будущем).

GitHub: github.com/Ekmor73

Благодарности
Fiber за удобный фреймворк.

Templ за мощный инструмент генерации шаблонов.

HTMX за упрощение динамического UI.

