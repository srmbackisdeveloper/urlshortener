FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o urlshortener cmd/server/main.go

FROM alpine:3.18

WORKDIR /app

COPY --from=builder /app/urlshortener .

EXPOSE 8080

CMD ["./urlshortener"]
