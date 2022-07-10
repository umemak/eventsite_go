FROM golang:1-alpine

RUN go install github.com/cosmtrek/air@latest

WORKDIR /app

CMD ["air"]
