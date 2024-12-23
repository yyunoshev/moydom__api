# moydom__api

## Требования
Перед тем как запустить проект, убедитесь, что у вас установлены следующие инструменты:

- Go — для сборки и запуска проекта.
- Swagger — для отображения документации.

## Установка
1. Клонируйте репозиторий
2. Установите зависимости:

В корне проекта выполните команду для скачивания зависимостей:
```bash
go mod tidy
```

В проекте используется файл .env для хранения переменных окружения, таких как настройки базы данных и секретный ключ для JWT.

Пример .env файла:

- SERVER_PORT=":8080"
- DATABASE_DB_URL="host=localhost user=iiunoshev password=iiunoshev dbname=moydom_dev port=5432 sslmode=disable"
- LOG_LEVEL="INFO"
- SECRET="auth-api-jwt-secret"
3. Сгенерируйте документацию:
```bash
swag init 
```
4. Запустите приложение:
```bash
go run cmd/main.go
```

Ссылка на документацию после запуска:
http://localhost:8080/swagger/index.html
