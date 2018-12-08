FROM golang:1.8

WORKDIR /go/src/github.com/frnd/schedule-hub/
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["schedule-hub"]

EXPOSE 8080