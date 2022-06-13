FROM golang:1.16-alpine

WORKDIR /usr/local/app
COPY go.mod go.sum /usr/local/app/

RUN go mod download
COPY . ./

RUN go build
CMD ./goshop
