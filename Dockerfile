FROM golang:1.19.3-alpine as dev
RUN mkdir -p /app

COPY /cmd/server.go /app/main.go
COPY .env /app/.env
COPY go.mod /app/go.mod
COPY go.sum /app/go.sum
WORKDIR /app
RUN mkdir -p api
COPY /api ./api
RUN go mod download
RUN go install github.com/cosmtrek/air@v1.29.0

CMD ["go", "run", "main.go"]
# CMD ["air", "-c", ".air.toml"]
