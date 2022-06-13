FROM golang:1.16-alpine

WORKDIR /usr/local/app
COPY go.mod go.sum /usr/local/app/

RUN go mod download
COPY . ./
RUN go build -o /docker-gs-ping
RUN make migrateUp

CMD ["bee", "run"]
