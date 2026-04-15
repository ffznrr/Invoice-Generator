# ===== BUILD STAGE =====
FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy

COPY . .

RUN go build -o main ./cmd/app/main.go

# ===== RUN STAGE =====
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main ./main

EXPOSE 8080

CMD ["./main"]