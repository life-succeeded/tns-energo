FROM golang:1.24 AS build

WORKDIR /build
COPY . .

# Скачиваем зависимости (кешируется в слое Docker)
RUN go mod download

# Копируем файлы, которые нужны в рантайме
RUN mkdir out && \
    mv .config/ out/

# Билдим гошечку в бинарник out/app
RUN go build -o out/app

FROM ubuntu:24.04

EXPOSE 8080

WORKDIR /app

COPY --from=build /build/out ./

ENTRYPOINT ["./app"]