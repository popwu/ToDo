FROM golang:1.23.1 AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o ./todo ./cmd/main.go
 
 
FROM alpine:latest AS runner
WORKDIR /app
COPY --from=builder /app/todo .
EXPOSE 8080
CMD ["./app/todo"]