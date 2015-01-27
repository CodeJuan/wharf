docker-bucket
===============

编译
====

代码 **clone** 到 `$GOPATH/src/githhub.com/dockercn` 目录下，然后执行以下命令

```
go get github.com/astaxie/beego
go get github.com/codegangsta/cli
go get github.com/siddontang/ledisdb/ledis
go get github.com/garyburd/redigo/redis                   <-好像还少一个库 忘了是什么了
go build
```
`TODO` 支持 **gopm** 编译程序

`TODO` 支持 **Dockerfile**

新建用户
=======

```
./docker-bucket account --action add --username docker --passwd docker --email bucket@docker.cn
```

对象 Key 规则
================

在 LedisDB 中对象的 Key 规则：

```
@Username // @用户名
#Organization // #组织名

@Username$Repository+ // 用户未加密公有仓库
#Organization$Repository+ // 组织未加密公有仓库

@Username$Repository- // 用户未加密私有仓库
#Organization$Repository- // 组织未加密私公有仓库

@Username$Repository-?Sign // 用户加密私有仓库  
#Organization$Repository-?Sign // 组织加密私有仓库

&Image+ //未加密，私有库和公有库未加密的 Image 共享
&Image-?Sign //加密，只有私有库有加密支持，每个 Image 根据加密签名不同，可能存有多份儿。

@Username$Repository*Template+(-) //  
#Organization$Repository*Template+(-) //

@Username$Repository!Job //  
#Organization$Repository!Job //
```

Bucket Conf
==========

```
[docker]
BasePath = /tmp/registry
StaticPath = files
Endpoints = 127.0.0.1
Version = 0.8.0
Config = prod
Standalone = true
OpenSignup = false
Encrypt = true

[ledisdb]
DataDir = /tmp/ledisdb
DB = 8

[email]
Host = smtp.exmail.qq.com
Port = 465
User = demo@docker.cn
Password = 123456

[log]
FilePath = /tmp/log
FileName = bucket-log
```

Nginx Conf
==========

Nginx 配置文件的示例，注意 **client_max_body_size** 对上传文件大小的限制。

```
upstream bucket_upstream {
  server 127.0.0.1:9911;
}

server {
  listen 80;
  server_name xxx.org;
  rewrite  ^/(.*)$  https://xxx.org/$1  permanent;
}

server {
  listen 443;

  server_name xxx.org;

  access_log /var/log/nginx/xxx.log;
  error_log /var/log/nginx/xxx-errror.log;

  ssl on;
  ssl_certificate /etc/nginx/ssl/xxx/x.crt;
  ssl_certificate_key /etc/nginx/ssl/xxx/x.key;

  client_max_body_size 1024m;
  chunked_transfer_encoding on;

  proxy_redirect     off;
  proxy_set_header   X-Real-IP $remote_addr;
  proxy_set_header   X-Forwarded-For $proxy_add_x_forwarded_for;
  proxy_set_header   X-Forwarded-Proto $scheme;
  proxy_set_header   Host $http_host;
  proxy_set_header   X-NginX-Proxy true;
  proxy_set_header   Connection "";
  proxy_http_version 1.1;

  location / {
    proxy_pass         http://bucket_upstream;
  }
}
```
