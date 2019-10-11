FROM alpine
LABEL younccat zzhbbdbbd@163.com

COPY . .

EXPOSE 3000

ENTRYPOINT ["/docker/scripts/entrypoint.sh"]
