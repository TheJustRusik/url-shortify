# Устанавливаем базовый образ для Go
FROM golang:1.24-alpine as builder

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем go.mod и go.sum для того, чтобы скачать зависимости
COPY go.mod go.sum ./
RUN go mod tidy

# Копируем исходный код приложения в контейнер
COPY . /app

# Собираем приложение
RUN go build -o myapp cmd/main.go

# Строим минимальный образ для запуска
FROM alpine:latest

# Копируем скомпилированный бинарник из предыдущего этапа
COPY --from=builder /app/myapp /usr/local/bin/myapp

# Указываем команду для запуска
CMD ["myapp"]
