FROM golang:1.20-alpine as builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go build -o app .
FROM alpine:latest

RUN apk --no-cache add ca-certificates
WORKDIR /app/
COPY --from=builder /app/app .
EXPOSE 8080

CMD ["./app"]
