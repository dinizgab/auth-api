FROM golang:1.24-alpine AS builder

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -ldflags="-s -w" -o /bin/auth ./cmd/main.go

FROM alpine:3.20  AS runtime

RUN apk add --no-cache ca-certificates && update-ca-certificates

WORKDIR /app
COPY --from=builder /bin/auth /app/auth

EXPOSE 8000

RUN adduser -D -u 10001 auth
USER auth

ENTRYPOINT ["/app/auth"]
