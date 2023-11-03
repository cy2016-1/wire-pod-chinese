gitee上传代码方法：

简易的命令行入门教程:

Git 全局设置:

`git config --global user.name "cv-robot"`
`git config --global user.email "xxx@xxx.com"`

创建 git 仓库:

mkdir wire-pod-chinese
cd wire-pod-chinese
git init 
touch README.md
git add README.md
git commit -m "first commit"
git remote add origin https://gitee.com/cv-robot/wire-pod-chinese.git
git push -u origin "master"`</u>`

已有仓库?

`cd existing_git_repo`
`git remote add origin https://gitee.com/cv-robot/wire-pod-chinese.git`
`git push -u origin "master"`
