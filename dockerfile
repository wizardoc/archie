FROM alpine
MAINTAINER younccat

COPY .                          .

EXPOSE 3000

CMD ./build/archie
