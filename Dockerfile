FROM golang:1.12

WORKDIR /go/src/app

ADD ./app .

RUN go build ./app

CMD ["./app"]