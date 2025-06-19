FROM golang:1.23-bullseye AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o ./bin/api ./cmd/api

RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.18.2/migrate.linux-amd64.tar.gz \
    | tar xvz && \
    chmod +x migrate


CMD ["/app/bin/api"]
EXPOSE 8080