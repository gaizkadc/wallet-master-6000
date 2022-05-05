FROM golang:1.17.2-alpine

RUN mkdir /app
WORKDIR /app

RUN apk add --update --no-cache

RUN go get -v github.com/cespare/reflex
EXPOSE 8080
ENTRYPOINT ["reflex", "-c", "docker/reflex.conf"]