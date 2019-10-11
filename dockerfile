FROM alpine
MAINTAINER younccat

WORKDIR /app

COPY .                          .

EXPOSE 3000

CMD ./build/archie
