FROM golang:1.20.0-alpine3.17

RUN mkdir /app

ADD . /app

WORKDIR /app

EXPOSE 8080

RUN go build -o main cmd/main.go

CMD ["/app/main"]
