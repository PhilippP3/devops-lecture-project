# Stage 1: Build stage
FROM golang:alpine AS builder

ARG SERVICE

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main ./$SERVICE/cmd/main.go


# Stage 2: Final stage
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .
EXPOSE 8080
CMD ["./main"]
