FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main ./main/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .

COPY --from=builder /app/.env ./

EXPOSE 8080

CMD ["sh", "-c", "ls -la /app && ./main"]

# CMD ["./main"]