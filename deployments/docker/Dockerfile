FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o go-backend-project ./cmd/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/go-backend-project .

EXPOSE 8081

CMD ["./go-backend-project"]