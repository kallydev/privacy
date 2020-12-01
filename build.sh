#!/bin/sh
export PATH=$PATH:/usr/local/go/bin
mkdir -p /opt/privacy/server/
mkdir -p /opt/privacy/website/
cd /tmp
git clone https://github.com/kallydev/privacy
cd /tmp/privacy/server
go build -o /opt/privacy/server/app main/main.go
chmod 0755 /opt/privacy/server/app
cd /tmp/privacy/website
yarn install && yarn build
cp -r /tmp/privacy/website/build /opt/privacy/website/build
rm -rf /tmp/privacy
rm -rf /root/go/pkg/*
rm -rf /usr/local/go/pkg/*
rm -rf /usr/local/share/.cache
