# server/Dockerfile
FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .
RUN go build -o cloudflareDnsCreator ./

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/cloudflareDnsCreator .

CMD ["./cloudflareDnsCreator"]
