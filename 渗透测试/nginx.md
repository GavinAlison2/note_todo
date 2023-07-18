# nginx access.log 记录详解

## 问题

- access.log记录的是什么


Nginx软件会把每个用户访问网站的日志信息记录到指定的日志文件里，供网站提供者分析用户的浏览行为等，此功能由ngx_http_log_module模块负责。对应的官方地址为：http://nginx.org/en/docs/http/ngx_http_log_module.html。

Nginx日志格式中默认的参数配置如下：

```text
 log_format  main  '$remote_addr - $remote_user［$time_local］"$request" '
'$status $body_bytes_sent "$http_referer" '
'"$http_user_agent" "$http_x_forwarded_for"'；
```

说明
```
$remote_addr    记录访问网站的客户端地址
$http_x_forward_for 当前端有代理服务器时，设置web节点记录客户端地址的位置，此参数生效的前提是代理服务器上也进行相关的x_forwarded_for设置
$remote_user    远程客户端用户名称
$time_local     记录访问时间与时区
$request        用户的http请求起始行信息
$status         http状态码，记录请求返回的状态，例如200,404,301等
$body_bytes_sents   服务器发送客户端的相应body字节数
$http_referer   记录此次请求时从哪个链接访问过来的，可以根据referer进行防盗链设置
$http_user_agent    记录客户端访问信息，例如：浏览器、手机客户端等


```