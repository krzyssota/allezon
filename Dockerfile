# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go get -u github.com/gin-gonic/gin
RUN go build -o allezon_server

EXPOSE 8090

CMD [ "./allezon_server" ]
