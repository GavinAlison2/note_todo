
# 抓取xyzq sec 这个app 后面的流量信息

## 抓取app后面的IP

通过 fiddler 抓取 xyzq sec 的包地址：  117.149.243.239

## 抓取app后面的域名

### 域名 s.thsi.cn  和 10jqka.com.cn 
域名 s.thsi.cn  

对应的网络服务商： 浙江贰贰网络有限公司
域名: 10jqka.com.cn
对应的网络服务商： 浙江贰贰网络有限公司

根据浙江贰贰网络有限公司的客服qq人说，对应的域名后面购买需要实名，
可以查查这两个域名的购买者 “s.thsi.cn”， “10jqka.com.cn”， 后面的个人实名信息
对应的网页地址: https://www.22.cn/

### ip  117.149.243.239

根据ip (117.149.243.239) 查询到对应的地址： 浙江省温州市
经纬度： 120.672111,28.000575
这个ip 地址后面的实名信息

对应的app, 请求地址： http://117.149.243.239:9099/app

## 抓取对应聊天的美洽软件与 对应美洽对接的企业ID

发现链接:
1. static.meiqia.com  GET /conversation/servable?ent_id=d96d402120638813502dd038e9910751 HTTP/1.1
2. static.meiqia.com  GET /fe-widget/v1.4.86.prod.20230712_46/552.js HTTP/1.1
3. new-api.meiqia.com GET /conversation/servable?ent_id=d96d402120638813502dd038e9910751 HTTP/1.1

域名： meiqia.com  
对应网络服务商： 美洽 https://meiqia.com/ 成都美洽网络科技有限公司  
    统一社会信用代码：9151010009889120X9

可加入美洽ios的开发群需求帮助，qq群，群1： 748720588， 群2： 515344099      
美洽客服电话： 400-0031-322     

获取人员信息：

```text
agentId=10413
accountType=1
agentName=四组
applyBuyAmt=1684642
authMsg=(null)
cardImg1=/uploadPath/cardImg/20230628/2d8ee267-2d73-481f-a9fa-d2598916d37f.jpg
cardImg2=/uploadPath/cardImg/20230628/c06b2a8b-7da2-4fe6-9e22-16e3b050f093.jpg
enableAmt=0.14
freezeAmt=207990
id=6952
idCard=362228199207014015
isActive=2
isAuthorize=0
isLock=False
isLogin=True
loginName=15711320691
nickName=新客
realName=黄勇
regIp=223.72.66.6
regTime=2023-06-28T12:08:37.000+0800
tradingAmount=0
userAmt=207756.8
```

接入美洽的请求：
```
GET new-api.meiqia.com/conversation/servable?ent_id=d96d402120638813502dd038e9910751 HTTP/1.1
Client
    Accept: application/json
    Accept-Encoding: gzip, deflate, br
    Accept-Language: zh-CN,zh-Hans;q=0.9
    User-Agent: Mozila/5.0(iPhone;CPU iPhone OS 16_2 like Mac OS X) AppleWebKit/605.1.15(KHTML, like Gecko) Mobile/15E148
Miscellaneous
    Referer: http://117.149.243.239:9099/
Security
    Origin: http://117.149.243.239:9099/
Transport
    Connection: keep-alive
    Host: new-api.meiqia.com
```

查看相关的美洽自定义聊天链接网址（https://meiqia.com/help/article/custom-domain/），下载对应的文件发现关键词ent_id
联合fiddler抓包接入美洽的请求发现： ent_id关键词， 我于2023-7-12 22:35 申请的企业id(ent_id) = e1f5030b0a040acb28937f8bec9cbe61 
推断出对应APP的接入美洽企业id是 **d96d402120638813502dd038e9910751**

### 主要的信息企业id(ent_id)= d96d402120638813502dd038e9910751

这个企业id需要联系 对应的美洽官方查询，咨询对方，对方回复：“信息无法给个人，警方需要我们会全力配合， 具体您可以让警方联系我们，可以线上咨询”

既然对方在使用美洽软件，那么是否可以查出使用方的ip地址
还有对应企业的申请信息，都可以查询到

登录页面

http://117.149.243.239:9099/app/pages/login/index

登录请求：

POST http://117.149.243.239:9099/apiApp/app/login

payload

{"loginName":"superadmin","loginPwd":"123456"}
