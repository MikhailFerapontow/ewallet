FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o ./build cmd/main.go

EXPOSE 3000

CMD ["sh", "-c", "sleep 30 && ./build"]