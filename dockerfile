# build app
FROM golang:alpine as builder

LABEL younccat zzhbbdbbd@163.com

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY=https://goproxy.cn

WORKDIR /build

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go get -u github.com/go-bindata/go-bindata/...
RUN go-bindata -o=assets/asset.go -ignore=".DS_Store|README.md|schema.go" -pkg=asset schema/...

RUN go build -o archie .

FROM alpine

COPY ./docker/scripts/wait-for-it.sh /

COPY ./templates /templates
COPY ./configs /configs
COPY --from=builder /build/assets /

COPY --from=builder /build/archie /

RUN apk update && apk add bash
RUN chmod 755 wait-for-it.sh