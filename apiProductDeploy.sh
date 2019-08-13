#!/usr/bin/env bash

publishVersion=$1
lumensFolder=`pwd`
startDeploy (){
    isVersion=`git tag | grep $publishVersion`
   #如果版本号没有就更新到最新版
    if [ "$isVersion" != "$publishVersion" ];then
            git  checkout master
            git  fetch --all
            git  reset  --hard origin/master
            git  pull
    else
        commitId=`git show $publishVersion -n 1|head  -n 1|awk '{print $2}'`
        git  reset --hard  $commitId
    fi

	    go build
        sudo systemctl restart qpgamego
}

#第一次部署
startFirtDeploy (){
sudo  sh -c "cat >/usr/lib/systemd/system/qpgamego.service<<EOF
[Unit]
Description=go  api project
After=network.target
[Service]
Type=simple
LimitNOFILE=65535
Environment=IRIS_MODE=release
Environment=ISAPISERVER=yes
Environment=ISTIMERSERVER=no
Environment=CURRENTPLATFORM=CKYX
Environment=ISCOMMONTIMERSERVER=yes
Restart=always
WorkingDirectory=/www/src/qpgame
ExecStart=/www/src/qpgame/qpgame
[Install]
WantedBy=multi-user.target
EOF"
sudo systemctl enable qpgamego.service
go get -u github.com/robfig/cron
go get -u github.com/Shopify/sarama
go get -u github.com/buger/jsonparser
go get -u github.com/dgrijalva/jwt-go
go get -u github.com/iris-contrib/middleware/jwt
go get -u github.com/go-redis/redis
go get -u github.com/go-sql-driver/mysql
go get -u github.com/iris-contrib/middleware/...
go get -u github.com/sirupsen/logrus
go get -u github.com/mojocn/base64Captcha
go get -u github.com/go-xorm/xorm
go get -u github.com/shopspring/decimal
go get -u github.com/mojocn/base64Captcha
go get -u github.com/wenzhenxi/gorsa


}
#需要传入版本参数
if [ ! $publishVersion ];then
    echo "请传入版本号参数 such as : v1.0.0"
elif  [ $publishVersion == "first_deploy" ];then
    startFirtDeploy
else
     startDeploy
fi


