FROM golang:1.25-alpine AS builder
WORKDIR /app
RUN apk add --no-cache build-base
COPY go.mod ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=1 GOOS=linux go build -o finance-api .

FROM alpine:3.21
WORKDIR /app
RUN apk add --no-cache ca-certificates sqlite
COPY --from=builder /app/finance-api ./finance-api
COPY --from=builder /app/.env.example ./.env.example
EXPOSE 8080
CMD ["./finance-api"]
