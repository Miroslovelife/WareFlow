# Этап сборки: используем образ Go для компиляции приложения
FROM golang:1.23.2 AS builder

# Устанавливаем необходимые пакеты для OR-Tools
RUN apt-get update && apt-get install -y \
    wget \
    unzip \
    && rm -rf /var/lib/apt/lists/*

# Скачиваем и распаковываем OR-Tools для Debian 11
RUN wget https://github.com/google/or-tools/releases/download/v9.11/or-tools_amd64_debian-11_cpp_v9.11.4210.tar.gz && \
    tar -xzf or-tools_amd64_debian-11_cpp_v9.11.4210.tar.gz -C /usr/local && \
    rm or-tools_amd64_debian-11_cpp_v9.11.4210.tar.gz

# Устанавливаем переменные окружения для OR-Tools
ENV LD_LIBRARY_PATH="/usr/local/or-tools/lib:${LD_LIBRARY_PATH}"
ENV CGO_LDFLAGS="-L/usr/local/or-tools/lib"
ENV CGO_CFLAGS="-I/usr/local/or-tools/include"

# Настраиваем рабочую директорию
WORKDIR /app

# Копируем go.mod и go.sum для установки зависимостей
COPY go.mod go.sum ./
RUN go mod download

# Копируем весь проект в контейнер
COPY . .

# Сборка приложения
RUN go build -o WareFlow ./cmd/main.go

# Финальный минимальный образ для запуска
FROM debian:buster-slim

# Устанавливаем переменные окружения для OR-Tools
ENV LD_LIBRARY_PATH="/usr/local/or-tools/lib:${LD_LIBRARY_PATH}"

# Переменные окружения для приложения
ENV MONGO_URI=mongodb://mongo:27017 \
    APP_ENV=production \
    PORT=8080

# Копируем собранное приложение и OR-Tools из этапа сборки
COPY --from=builder /app/WareFlow /app/WareFlow
COPY --from=builder /usr/local/or-tools /usr/local/or-tools

# Открываем порт для приложения
EXPOSE 8082

# Запускаем приложение
CMD ["/app/WareFlow"]
