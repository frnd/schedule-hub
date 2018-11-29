FROM golang:latest

RUN mkdir -p /app
WORKDIR /app

ADD . /app
RUN go get -t -v ./...
RUN go build ./main.go

EXPOSE 8080

CMD ["./app"]