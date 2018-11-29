FROM golang:latest

RUN mkdir -p /app
WORKDIR /app

ADD . /app
RUN go get -t -v ./... \
 && go build -v

EXPOSE 8080

CMD ["./schedule-hub"]