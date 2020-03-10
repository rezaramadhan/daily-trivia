FROM golang:1.14 as builder

COPY . /app

WORKDIR /app

RUN go mod download

RUN go build get-trivia.go

ENTRYPOINT ["./get-trivia"]
