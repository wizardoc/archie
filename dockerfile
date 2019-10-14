FROM alpine
LABEL younccat zzhbbdbbd@163.com

COPY . .

ADD ./docker/scripts/wait-for-it.sh /usr/local/bin/

EXPOSE 3000

ENTRYPOINT ["/docker/scripts/entrypoint.sh"]
