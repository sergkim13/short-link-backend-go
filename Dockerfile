
FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN GOOS=linux go build -o sl-backend cmd/main.go

FROM alpine:latest
WORKDIR /
COPY --from=builder /app/sl-backend /sl-backend
EXPOSE 8080
CMD ["/sl-backend"]
