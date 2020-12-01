FROM node:15-buster

RUN apt update && apt install software-properties-common -y && apt clean && rm -rf /var/lib/apt/lists/*

RUN apt-key adv --keyserver keyserver.ubuntu.com --recv-keys CC86BB64

RUN add-apt-repository "deb http://ppa.launchpad.net/rmescandon/yq/ubuntu focal main" -y

RUN apt update && apt install sqlite python git moreutils yq -y && apt clean && rm -rf /var/lib/apt/lists/*

WORKDIR /tmp

ADD https://golang.org/dl/go1.15.5.linux-amd64.tar.gz /tmp/go1.15.5.linux-amd64.tar.gz

RUN tar -C /usr/local -xzf go1.15.5.linux-amd64.tar.gz && echo 'export PATH=$PATH:/usr/local/go/bin' >> /etc/profile

WORKDIR /opt

RUN git clone https://github.com/kallydev/privacy

WORKDIR /opt/privacy/server

RUN export PATH=$PATH:/usr/local/go/bin && go build -o app main/main.go

WORKDIR /opt/privacy/website

RUN yarn install && yarn build

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

WORKDIR /opt/privacy

ENTRYPOINT [ "/opt/privacy/entrypoint.sh" ]
