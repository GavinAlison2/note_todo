## 原理

内网穿透由三部分组成：Control端（A）、bridge&server端（B）、agent端（C）。Control与server通过bridge进行桥接。

系统：centos7
A：内网ip 10.9.102.33（vpn）
B：公网ip 123.206.66.166  内网ip 10.9.102.28（vpn）
C：内网ip 192.168.228.129（无外网ip，能访问外网）

## 需求

使A能通过B连接到（ssh）C机器的22端口

A --> C 通过 B 转发
A 无法 ping 通 B


## 实现

使用工具：goproxy 官方地址：https://github.com/snail007/goproxy

在 B 和 C 搭建一个内网穿透工具，实现 A 访问B 端口， 转发端口到C，从而间接实现 A 访问C 的服务或者tcp

## 部署

### 安装goproxy

分别在B和C上安装goproxy工具

```
自动安装：
curl -L https://raw.githubusercontent.com/snail007/goproxy/master/install_auto.sh | bash  

手动安装(安装前自己创建部署目录)：
mkdir /home/proxy
cd /home/proxy
##下载守护进程monexec
wget https://github.com/reddec/monexec/releases/download/v0.1.1/monexec_0.1.1_linux_amd64.tar.gz

##下载proxy
wget https://github.com/snail007/goproxy/releases/download/v3.7/proxy-linux-amd64.tar.gz

##下载自动安装脚本：
wget https://raw.githubusercontent.com/snail007/goproxy/master/install.sh

##安装
/bin/bash install.sh

```

### 配置内网穿透

#### B机器

```
##创建proxy的公钥和私钥文件
proxy keygen -C proxy

##建立端口映射
proxy tbridge -p ":33080" -C proxy.crt -K proxy.key --daemon
proxy tserver -r ":2202@:22" -P "127.0.0.1:33080" -C proxy.crt -K proxy.key --daemon
//参数的详细解释看官方文档

完成！
```

#### C机器

```
##拷贝B机器的公钥和秘钥文件：
scp 123.206.66.166:/usr/local/src/goproxy/proxy.key /home/proxy/

##建立映射
proxy tclient -P "123.206.66.166:33080" -C proxy.crt -K proxy.key

Done！
```

## 连接

A 机器 通过B 连接A: 22，

```
##远程连接测试
ssh -p 2202 root@123.206.66.166
//-p：指定端口（默认是22端口）

password：输入C机器的密码即可！
```
