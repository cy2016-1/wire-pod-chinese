


简易的命令行入门教程:

Git 全局设置:

git config --global user.name "cv-robot"
git config --global user.email "xxx@xxx.com"

创建 git 仓库:

mkdir wire-pod-chinese
cd wire-pod-chinese
git init 
touch README.md
git add README.md
git commit -m "first commit"
git remote add origin https://gitee.com/cv-robot/wire-pod-chinese.git
git push -u origin "master"

已有仓库?

cd existing_git_repo
git remote add origin https://gitee.com/cv-robot/wire-pod-chinese.git
git push -u origin "master"




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
