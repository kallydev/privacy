FROM node:15-alpine3.12

RUN echo "http://dl-cdn.alpinelinux.org/alpine/edge/community" >> /etc/apk/repositories
RUN apk add --update --no-cache sqlite python3 git moreutils yq gcc make libc-dev wget

# 修复alpine的golang支持
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2

WORKDIR /tmp

RUN wget https://golang.org/dl/go1.15.5.linux-amd64.tar.gz && tar -C /usr/local -xzf go1.15.5.linux-amd64.tar.gz && echo 'export PATH=$PATH:/usr/local/go/bin' >> /etc/profile && rm -f /tmp/go1.15.5.linux-amd64.tar.gz

ADD build.sh /tmp/build.sh
RUN chmod +x /tmp/build.sh && /tmp/build.sh

WORKDIR /opt/privacy

ADD scripts/ /opt/privacy/scripts/
ADD config.yaml /opt/privacy/config.yaml
ADD entrypoint.sh /opt/privacy/entrypoint.sh
RUN chmod +x /opt/privacy/entrypoint.sh

VOLUME [ "/opt/privacy/database", "/opt/privacy/source" ]

ENV qq=true
ENV jd=true
ENV sf=false
ENV wb=false

ENV host=0.0.0.0
ENV port=80

ENV mask=true

EXPOSE 80

ENTRYPOINT [ "/opt/privacy/entrypoint.sh" ]
