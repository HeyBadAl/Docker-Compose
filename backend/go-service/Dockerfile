FROM golang:1.21-alpine3.17

WORKDIR /go/src/app

COPY . .

RUN go build -o main .

EXPOSE 8080

CMD ["./main"]

