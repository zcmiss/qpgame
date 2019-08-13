### qpgameApi go项目说明


# 一、代码结构说明

#### 1. 代码结构
#### /admin 后台代码
##### /admin/common 后台通用代码
##### /admin/controllers 后台控制器
##### /admin/models 后台模型代码
##### /admin/jobs 后台执行的定时任务管理系统
##### /admin/validations 后台校验器代码, 用于添加/修改操作之前的验证
##### /admin/websockets 用于WebSocket的代码
#### /app 前台代码
##### /app/frontEndControllers 前台控制器
##### /app/frontEndControllers/pay 前台关于支付的代码
##### /app/fund 通用于前于资金操作的代码
#### /assets 静态资源文件
#### /common 前后、后台通用文件、函数、库
##### /common/log 日志相关函数
##### /common/mvc 通用 model-controller-validation 相关函数
##### /common/timer 定时器相关
##### /common/utils 通用于工具函数
#### /config 项目配置相关文件
###### /config/apiStatusCode.go 关于api全局常量的定义
###### /config/fundConfig.go 关于资金相关的常量定义
###### /config/gameConfig 关于游戏下样的常量定义
###### /config/ip2region.db ip转换数据库
###### /config/iris.yml 框架本身的相关配置
###### /config/irisGlobal.go  框架相关全局变量定义
###### /config/tpPayConfig.go 支付相关全局变量定义
#### /doc 项目文件、需求说明
#### /middlewares 系统中间件，包括前台、后台
###### /middlewares/beforeMiddleware.go 所有路由之前的判断
###### /middlewares/corsMiddleware.go CROS相关的判断
###### /middlewares/jwtAuthenticate.go JWT相关判断
#### /models Go的orm库
##### /models/mainxorm XORM库
###### /models/mysql MYSQL库
###### /models/pool 连接池
###### /models/xorm 程序ORM
#### /node_modules 用于apidoc生成的node库, go开发用不到
#### /ramcache 内存缓存，常驻内存，前台、后台都有用到
#### /routers 路由设置，包括前台和后台
###### /routers/admin.go 后台路由设置
###### /routers/web.go 前台路由设置
#### /views 视图文件，目前项目没有用到
#### /main.go 程序入口文件

#### 2. 程序运行流程
```
    客户端(访问)
        -> 路由判断(/routers) 
        -> 中间件处理(/middlewares)
        -> 控制器/方法(/admin/controllers或/app/frontEndControllers)
        -> 模型相关方法/通行方法/内存缓存处理等
        -> 控制器/方法(同上)
        -> 客户端(返回)
```


# 二、相关工具、框架、库

#### 1.框架安装
`go get -u github.com/kataras/iris`

#### 2.安装热编译工具可以提高开发效率
`go get -u github.com/silenceper/gowatch`
* 在linux上的当前目录执行gowatch启动
* 在gowatch.yml中已配置环境变量，包括是否是线上环境
* 是否是接口服务器还是后台管理接口服务器,开发的时候不要
* 修改代码，修改这个gowatch.yml配置文件即可

#### 4.安装kafka(window环境不一样,到linux上测试)
* 安装必要的库: 

`yum install -y gcc gcc-c++ pcre-devel zlib-devel`
* 下载并安装rdkafka

``git clone https://github.com/edenhill/librdkafka.git
cd librdkafka
./configure
make && make install``
##### 4.2 安装sarama
`go get -u github.com/Shopify/sarama`

#### 6.安装json字节操作库
`go get -u github.com/buger/jsonparser`

#### 7.下载jwt
``go get -u github.com/dgrijalva/jwt-go
go get -u github.com/iris-contrib/middleware/jwt``

#### 8.下载redis
`go get -u github.com/go-redis/redis`

#### 9.下载安装mysql
`go get -u github.com/go-sql-driver/mysql`

#### 10.定时器安装
`go get -u github.com/robfig/cron`

* [定时器参考文档](https://www.jianshu.com/p/626acb9549b1)

#### 11.安装iris中间件
`go get -u github.com/iris-contrib/middleware/...`

#### 12.安装日志工具
`go get -u github.com/sirupsen/logrus`

#### 13.安装xrom数据库映射工具
`go get -u github.com/go-xorm/xorm`

#### 14.安装decimal数据库映射工具
`go get -u github.com/shopspring/decimal`

#### 15.安装下验证码工具
`go get -u github.com/mojocn/base64Captcha`

#### 16.安装FTP工具
`go get -u github.com/jlaffaye/ftp`

#### 17选装 xorm生成工具
`go get -u github.com/go-xorm/cmd/xorm`
* 进入gopath\src\github.com\go-xorm\cmd\xorm编译执行go build
* 命令如下
* ```xorm reverse mysql "ckgame:+g*=!@tcp(52.78.44.55:3306)/ckgame?charset=utf8&multiStatements=true" templates/goxorm```
* 在当前的models目录下生成结构体映射

#### 18.安装二维码工具
`go get -u github.com/skip2/go-qrcode`

#### 19.安装ip转换库
`go get -u github.com/lionsoul2014/ip2region/binding/golang`
* [IP转换参考文档](https://github.com/lionsoul2014/ip2region/tree/master/binding/golang)


#### 使用说明: 根据错误等级调用特定的方法
`util.Log.WithFields(logrus.Fielsd{}).Info() `
