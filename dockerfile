FROM alpine
MAINTAINER younccat

WORKDIR /archie
COPY .                          /archie

EXPOSE 3000

ENTRYPOINT ["/archie/build/archie"]
