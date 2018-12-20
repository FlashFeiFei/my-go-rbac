# 项目代号： 兔子(tuzi)

不懂编程的同学听我演示一遍就可以拥有一个企业官网了！你只需安装好Golang环境，安装教程网站一搜很多。现在v1版本支持:

1. 基本的登录(需要验证码), cookie记住登录功能
2. 权限认证功能(基于RBAC), 有友好的后台操作界面, 可以进行角色性的功能授权, 用户管理
3. 支持基本的文章(可置顶,指定排序), 富文本编辑, 有回收站功能. 相册功能, 可上传文件, 类似文章功能
4. 前端支持轮转图, 首页配置灵活, 可配置统计代码, 跟帖代码等
5. 初始化数据将使用知乎回答, 待做!!

见 [图片](/doc/example.md)

![兔子](/tuzi.png)

## 一.使用说明

使用只需拉下库:

```shell
go get -v github.com/hunterhug/rabbit
```

或者:

```
git clone https://www.github.com/hunterhug/rabbit
mkdir -p %GOPATH%/src/github.com/hunterhug
mv rabbit %GOPATH%/src/github.com/hunterhug
```

编译程序

```shell
go build
```

启动前请配置`conf/app.conf`中的数据库（安装Mysql请自行百度）

```
# 生产环境可改为prod
runmode = dev

[dev]
httpport = 8080
db_host = 127.0.0.1
db_port = 3306
db_user = root
db_pass = root
db_name = tuzi
db_type = mysql

[prod]
EnableGzip = true
httpport = 80
db_host = 127.0.0.1
db_port = 3306
db_user = root
db_pass = root
db_name = tuzi
db_type = mysql
```

初始化数据库

```shell
./rabbit -db=1
```

运行程序,调试建议使用beego官方工具`bee run`

```shell
./rabbit
```

这时，你可以打开`http://127.0.0.1:8080`, 进入后台编辑网站：`http://127.0.0.1:8080/public/login`, 账户`admin`, 密码：`admin`

如果上传文件出错，请在本项目新建file文件夹并赋予权限。

```
# if in linux
mkdir file
chmod 777 file
```

## 二.开发流程

### a.文件结构

```shell
    ----conf 配置文件夹

        ----app.conf 		应用配置文件
        ----local_**.ini 	国际化文件

    ----controllers 控制器
        ----admin	后台控制器
            ----blog 博客模块
            ----rbac 权限模块
        ----home 	前台控制器
        ----rbac.go 路由权限过滤器入口

    -----lib 公共库
    -----file 上传文件保存地址
    -----models ORM模型
        ----admin RBAC主要数据库
            ----AdminInit.go 数据默认填充地
        ----blog  博客主要数据库

    ----routers 路由
    ----static  静态文件
        ---admin 后台js/css勿改
        ---home 前台UI美观第三方js/css
            ---amazi  妹紙UI
            ---boostrap 最牛逼的界面UI
         ---tool 公用第三方js
         ---diy 自己的js/css
    ----views	视图
        ----admin 	后台视图
            ----default 默认主题
        ----home 	前台视图
            ----default 默认主题

    ----front 前端测试文件夹
        ---data 模拟的JSON数据

    ---help  帮助脚本等
        --- init.sql 提供的示例数据库
        --- ngnix-tuzi.conf Nginx配置

    ----doc 说明文档
    ----test 测试文件夹
```

目前实现了基本的RBAC模块和博客模块（文章和相册）,附带Dashboard后端,UI基本框架形成依靠此项目可敏捷开发.

1. 基于角色的访问控制（Role-Based Access Control）作为传统访问控制
2. Amaze UI v2.7.0（部分后台）和jQuery EasyUI v1.4.2（后台表格CRUD）、Bootstrap v3.3.5（前台）混合
3. 准备采用Vue.js v2.2.6 前后端完全分离（Maybe）,后台写死很笨拙但是对前台开放友好的REST JSON API这样可离线测试前端.
ajax调用JSON时请注意跨域问题(见front文件夹),这样的好处是将渲染视图的工作交给用户的浏览器端.(可不选择)

### b.约定

1. RBAC权限相关的models统一放在admin文件夹,其他都放在home文件夹.前台控制相关的controllers统一放在home文件夹,其他都放在admin文件夹.URL router统一`M/C/A`方式,该正则url需要验证权限如rbac/public/index（三个参数）,其他如public/index不验证.
2. 登录说明：登陆过的用户只能注销后登录，支持定义cookie登录.进入后台时验证session,session不存在则验证cookie.如果用户未被冻结,增加session,同时更改用户登录时间、登录IP等.cookie与登录IP绑定（防止cookie劫持）.
3. 系统时间默认数据库本地时间为东八区北京时间.
4. 后台模板在`views/admin`前台模板在`views/home`子文件夹为主题默认主题为default
5. 所有配置在conf文件夹`conf/app.conf`支持国际化
6. 数据库数据填充在`models/*/*Init.go`中定义， 我准备将所有中文变成英文
7. 各种前端文件全部放在`static`中
8. 前台首页配置（可动态调整首页待解释）

```
{
	    "1":{"name":"About","limit":6},
        "2":{"name":"News","limit":6},
        "3":{"name":"Lifes","limit":6},
        "4":{"name":"Production","limit":6},
        "5":{"name":"Flower","limit":6},
        "6":{"name":"TeaCup","limit":6}
}
```

### c.增加路由和权限

每次在`models/admin/AdminInit.go`增加路由权限请执行

```
./rabbit -rbac=1
```

调试请使用`bee run`

## 三.Nginx架站（可选）

请百度安装nginx，Ubuntu用户可以`sudo apt-get install nginx`, 功能：反向代理，将8080端口映射到一个域名的80端口上，你只需A记录到该域名即可。

进入`/etc/nginx/conf.d`, 放入该项目下`doc/sh/ngnix-tuzi.conf`文件, 下面的配置`ngnix-tuzi.conf`,`server_name`为域名,`access_log`为日志路径（要手动建文件夹）

```shell
server{
        listen 80;
        server_name tuzi.lenggirl.com;
        charset utf-8;
        access_log /data/logs/nginx/tuzi.lenggirl.com.log;
        #error_log /data/logs/nginx/tuzi.lenggirl.com.err;
        location / {
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header Host $http_host;
        proxy_redirect off;
        proxy_pass http://localhost:8080;
	    proxy_set_header X-Real-Ip $remote_addr;
        }

}
```

然后:

```
nginx -t
nginx -s reload
curl tuzi.lenggirl.com
```

## 四.环境配置

安装Golang环境请百度

## 五.精彩演示

见 [图片](/doc/example.md)

如果你觉得项目帮助到你,欢迎请我喝杯咖啡

微信
![微信](https://raw.githubusercontent.com/hunterhug/hunterhug.github.io/master/static/jpg/wei.png)

支付宝
![支付宝](https://raw.githubusercontent.com/hunterhug/hunterhug.github.io/master/static/jpg/ali.png)


问题咨询请发邮件:gdccmcm14@live.com.

# LICENSE
    
欢迎加功能(PR/issues),请遵循Apache License协议(即可随意使用但每个文件下都需加此申明）

```
Copyright 2017 by rabbit author: gdccmcm14@live.com.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License
```

