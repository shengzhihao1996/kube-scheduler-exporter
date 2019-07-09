FROM alpine

COPY bin/app /app/app

COPY kubectl /usr/local/bin

COPY config /etc/config

RUN echo -e "https://mirrors.aliyun.com/alpine/v3.8/main\nhttps://mirrors.aliyun.com/alpine/v3.8/community" > /etc/apk/repositories \
    && apk add --no-cache curl net-tools \
    && chmod +x /app/app  /usr/local/bin/kubectl

WORKDIR /app

USER root

CMD ["./app"]
