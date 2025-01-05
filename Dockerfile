FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go build -o build/main cmd/main.go


FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/build/* .
COPY --from=builder /app/config.yaml .

EXPOSE 8080

# Command to run the application
CMD ["./main"]