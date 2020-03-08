FROM golang:1.13.4-alpine3.10 AS build

RUN apk update && apk add --virtual build-dependencies build-base --no-cache curl ca-certificates git dep gcc

ENV GOROOT=/usr/local/go \
  GOPATH=/app

ADD . /app/src

WORKDIR /app/src

RUN go build -o bin/server main.go

FROM alpine:3.10
WORKDIR /app

COPY --from=build /app/src/bin/ /app/

RUN ls

ENTRYPOINT ["/app/server"]
