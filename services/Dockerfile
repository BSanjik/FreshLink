# Этап 1: билд
FROM golang:1.24.2 AS builder

WORKDIR /app

# Зависимости
COPY go.mod go.sum ./
RUN go mod download

# Копируем весь код
COPY . .

# Сборка
RUN go build -o freshlink ./cmd

# Этап 2: финальный образ
FROM alpine:latest

WORKDIR /app

# Установка сертификатов
RUN apk --no-cache add ca-certificates

# Копируем бинарник из builder
COPY --from=builder /app/freshlink .

# Используем порт
EXPOSE 8080

# Старт приложения
CMD ["./freshlink"]