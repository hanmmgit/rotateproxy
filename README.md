# rotateproxy

利用fofa搜索socks5开放代理进行代理池轮切的工具

## 特性

- 支持认证
- 代理列表优选策略：随机、或延时前多少位随机

## 帮助

```shell
> .\rotateproxy.exe -h
Usage of rotateproxy.exe:
  -email string
        email address
  -l string
        listen address (default ":8899")
  -page int
        the page count you want to crawl (default 5)
  -pass string
        authentication password
  -region int
        0: all 1: cannot bypass gfw 2: bypass gfw
  -rule string
        search rule (default "protocol==\"socks5\" && \"Version:5 Method:No Authentication(0x00)\" && after=\"2021-08-01\" && country=\"CN\"")
  -token string
        token
  -user string
        authentication username
```

## 安装

```shell
go get -u github.com/akkuman/rotateproxy/cmd/...
```

或者到[发布页面](https://github.com/akkuman/rotateproxy/releases/latest)进行下载

### 安装为linux服务（感谢 [@Rvn0xsy](https://github.com/Rvn0xsy) 提供 [PR](https://github.com/akkuman/rotateproxy/pull/4)）

1. 下载相关文件

```bash
curl -o /usr/local/bin/rotateproxy 'https://github.com/akkuman/rotateproxy/releases/latest/download/rotateproxy-linux-amd64'
curl -o /usr/lib/systemd/system/rotateproxy.service 'https://raw.githubusercontent.com/akkuman/rotateproxy/master/rotateproxy.service.example'
```

2. 查看 [rotateproxy.service.example](./rotateproxy.service.example) 文件示例，将 `/usr/lib/systemd/system/rotateproxy.service` 文件中的 `ExecStart` 的命令替换为你自己的命令
3. 启动服务

```bash
# 开启服务
service rotateproxy start
# 关闭服务
service rotateproxy stop
# 重启服务
service rotateproxy restart
# 设置开机自启动
systemctl enable rotateproxy.service
```


## 效果展示

![](./pics/curl-run.jpg)