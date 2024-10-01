FROM golang:1.20 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .

# Pass environment variables (optional)
# ENV DISCORD_TOKEN=<your-token-here> or pass it during runtime.

ENTRYPOINT ["./main"]
