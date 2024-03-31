FROM golang:1.22.1-alpine3.19
RUN apk update && apk add git && apk add bash && apk add curl
COPY . /app
WORKDIR /app
