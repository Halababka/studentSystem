# syntax=docker/dockerfile:1

FROM golang:1.18-alpine
WORKDIR /api
COPY go.mod ./
RUN go mod download
COPY *.go ./

RUN go get
RUN go build -o /stusy-api

EXPOSE 8080

CMD [ "/stusy-api" ]
