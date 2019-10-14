FROM alpine
LABEL younccat zzhbbdbbd@163.com

COPY . .

ADD ./docker/scripts/wait-for-it.sh /usr/local/bin/

RUN echo "https://mirror.tuna.tsinghua.edu.cn/alpine/v3.4/main/" > /etc/apk/repositories

RUN apk update \
        && apk upgrade \
        && apk add --no-cache bash \
        bash-doc \
        bash-completion \
        && rm -rf /var/cache/apk/* \
        && /bin/bash
EXPOSE 3000

ENTRYPOINT ["/docker/scripts/entrypoint.sh"]
