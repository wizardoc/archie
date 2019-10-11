FROM alpine
MAINTAINER younccat

COPY .                          .

EXPOSE 3000

WORKDIR /build

CMD ./archie
