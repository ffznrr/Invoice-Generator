FROM golang:1.22-alpine

WORKDIR /app

RUN apk add --no-cache git


RUN go install github.com/cosmtrek/air@v1.50.0


COPY go.mod go.sum ./
RUN go mod tidy


COPY . .

EXPOSE 8080

CMD ["air"]

# for Prod
# # ===== BUILD STAGE =====
# FROM golang:1.22-alpine AS builder

# WORKDIR /app

# COPY go.mod go.sum ./
# RUN go mod tidy

# COPY . .

# RUN go build -o main ./cmd/app/main.go

# # ===== RUN STAGE =====
# FROM alpine:latest

# WORKDIR /app

# COPY --from=builder /app/main .

# EXPOSE 8080

# CMD ["./main"]