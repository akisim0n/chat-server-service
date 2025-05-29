FROM golang:1.24-alpine3.21 AS builder

COPY . /github.com/akisim0n/chat-server-service/
WORKDIR /github.com/akisim0n/chat-server-service/

RUN go mod download
RUN go build -o ./bin/chat_server cmd/server/main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /github.com/akisim0n/chat-server-service/application/bin/chat_server .

CMD ["./chat_server"]

