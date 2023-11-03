

## 运行前修改chipper/start.sh文件配置密钥和python路径：

需要在百度云控制台和迅飞大模型控制台申请对应的免费api key：

`export BAIDU_APIKEY="xxx"`

`export BADIU_APISECRET="xxx"`



`export XUNFEI_APPID="xxx"`

`export XUNFEI_APISECRET="xxx"`

`export XUNFEI_APIKEY="xxx"`



`export PYTHON_CMD="python3"`

`export PYTHON_CHDIR="xxx/wire-pod-chinese/python"`



## 安装方法：

参考官方的安装过程，出了问题可以到官方github的issues搜索解决办法：

The installation guide exists on the wiki: [Installation guide](https://github.com/kercre123/wire-pod/wiki/Installation)



#### go仓库加速：

实际上安装只需要输入sudo ./setup.sh，只是国内很可能无法成功下载需要的包，

解决方法：

`go env -w GO111MODULE=on`

`go env -w GOPROXY=https://goproxy.cn,direct #这个命令会让下载github.com的go包加速`



#### 如何让局域网连接到服务器escapepod.loacal:

运行之前需要修改自己的用户名为escapepod，但是并不推荐，可以用如下方法让局域网也可以ping通escapepod：

`sudo apt-get install avahi-daemon`

`sudo gedit /etc/avahi/avahi-daemon.conf`

`找到#host-name=foo这一行，取消注释并将其修改为：`

`host-name=escapepod`

`保存文件并退出编辑器。`

`sudo service avahi-daemon restart`



#### 系统要求：

官方的方法里写了可以在windows和linux两种系统下安装，我只尝试了ubuntu电脑和香橙派。

如果不想一直开着电脑，可以在树莓派或者香橙派上部署本项目，并添加开机自启动服务，这样就可以完美实现和vector对话了。



#### python环境安装：

python如果报错，根据报错信息用pip或者apt安装缺少的包即可。



#### 技术支持：

有问题可以先在gitee提issues，如果安装实在有困难，也可以联系我的weichat：nqluowei



## 运行方法：

sudo ./chipper/start.sh



## 官方的原始README：


# wire-pod

`wire-pod` is fully-featured server software for the Anki (now Digital Dream Labs) [Vector](https://web.archive.org/web/20190417120536if_/https://www.anki.com/en-us/vector) robot. It was created thanks to Digital Dream Labs' [open-sourced code](https://github.com/digital-dream-labs/chipper).

It allows voice commands to work with any Vector 1.0 or 2.0 for no fee, including regular production robots.

## Installation

The installation guide exists on the wiki: [Installation guide](https://github.com/kercre123/wire-pod/wiki/Installation)

## Wiki

Check out the [wiki](https://github.com/kercre123/wire-pod/wiki) for more information on what wire-pod is, a guide on how to install wire-pod, troubleshooting, how to develop for it, and for some generally helpful tips.

## Donate

If you want to :P

[![Donate](https://img.shields.io/badge/Donate-PayPal-green.svg)](https://www.paypal.com/donate/?business=53VQ3Q95TD2M6&no_recurring=0&currency_code=USD)

## Credits

- [Digital Dream Labs](https://github.com/digital-dream-labs) for saving Vector and for open sourcing chipper which made this possible
- [dietb](https://github.com/dietb) for rewriting chipper and giving tips
- [fforchino](https://github.com/fforchino) for adding many features such as localization and multilanguage, and for helping out
- [xanathon](https://github.com/xanathon) for the publicity and web interface help
- Anyone who has opened an issue and/or created a pull request for wire-pod
