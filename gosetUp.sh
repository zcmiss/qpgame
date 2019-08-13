#!/usr/bin/env bash
yum -y update
yum -y install wget vim git
cd /usr/local
wget https://dl.google.com/go/go1.12.1.linux-amd64.tar.gz
tar -zxvf go1.12.1.linux-amd64.tar.gz
rm -rf go1.12.1.linux-amd64.tar.gz
export GOROOT=/usr/local/go
cd /
mkdir www
sh -c "cat >> ~/.bashrc <<EOF
export GOROOT=/usr/local/go
export GOPATH=/www
export PATH=\$GOPATH/bin:\$GOROOT/bin:\$PATH
EOF"

sh -c "cat >> /etc/profile <<EOF
export GOROOT=/usr/local/go
export GOPATH=/www
export PATH=\$GOPATH/bin:\$GOROOT/bin:\$PATH
EOF"
# 设置时区
timedatectl set-timezone Asia/Shanghai
source /root/.bashrc

source /etc/profile

go version
go env
go get -u github.com/kataras/iris