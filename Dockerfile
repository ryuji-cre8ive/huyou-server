# goバージョン
FROM golang:1.18.0-alpine

RUN apk update && apk add git && apk add alpine-sdk

ENV APP_ROOT /app

RUN mkdir $APP_ROOT

WORKDIR $APP_ROOT

ENV CGO_ENABLED=1 \
  GOOS=linux \
  GOARCH=amd64 \
  GO111MODULE=on

COPY ./go.mod $APP_ROOT/go.mod
COPY ./go.sum $APP_ROOT/go.sum

COPY . $APP_ROOT

RUN go mod download

EXPOSE 8080

CMD [ "go", "run", "server.go" ]