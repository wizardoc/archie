FROM golang:latest
MAINTAINER younccat

WORKDIR /app

COPY . /app

EXPOSE 3000

ENTRYPOINT ["/app/main"]
