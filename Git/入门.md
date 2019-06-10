****什么是Git**?**

一、概述

1 基本概念：

Git 是一个开源的分布式版本控制系统，用于敏捷高效地处理任何或小或大的项目。

Git 是 Linus Torvalds 为了帮助管理 Linux 内核开发而开发的一个开放源码的版本控制软件。

Git 与常用的版本控制工具 CVS, Subversion 等不同，它采用了分布式版本库的方式，不必服务器端软件支持。

2 Git 与 SVN 区别

Git 不仅仅是个版本控制系统，它也是个内容管理系统(CMS)，工作管理系统等。

如果你是一个具有使用 SVN 背景的人，你需要做一定的思想转换，来适应 Git 提供的一些概念和特征。

Git 与 SVN 区别点：

1、Git 是分布式的，SVN 不是：这是 Git 和其它非分布式的版本控制系统，例如 SVN，CVS 等，最核心的区别。

2、Git 把内容按元数据方式存储，而 SVN 是按文件：所有的资源控制系统都是把文件的元信息隐藏在一个类似 .svn、.cvs 等的文件夹里。

3、Git 分支和 SVN 的分支不同：分支在 SVN 中一点都不特别，其实它就是版本库中的另外一个目录。

4、Git 没有一个全局的版本号，而 SVN 有：目前为止这是跟 SVN 相比 Git 缺少的最大的一个特征。

5、Git 的内容完整性要优于 SVN：Git 的内容存储使用的是 SHA-1 哈希算法。这能确保代码内容的完整性，确保在遇到磁盘故障和网络问题时降低对版本库的破坏。


3 基本命令

git init //初始化本地git环境

git clone XXX//克隆一份代码到本地仓库

git pull //把远程库的代码更新到工作台

git pull --rebase origin master //强制把远程库的代码跟新到当前分支上面

git fetch //把远程库的代码更新到本地库

git add . //把本地的修改加到stage中

git commit -m 'comments here' //把stage中的修改提交到本地库

git push //把本地库的修改提交到远程库中

git branch -r/-a //查看远程分支/全部分支

git checkout master/branch //切换到某个分支

git checkout -b test //新建test分支

git branch -d test //删除test分支 -D强制删除

git merge master //假设当前在test分支上面，把master分支上的修改同步到test分支上

git merge tool //调用merge工具

git stash //把未完成的修改缓存到栈容器中

git stash list //查看所有的缓存

git stash pop //恢复本地分支到缓存状态

git blame someFile //查看某个文件的每一行的修改记录（）谁在什么时候修改的）

git status //查看当前分支有哪些修改

git log //查看当前分支上面的日志信息

git diff //查看当前没有add的内容

git diff --cache //查看已经add但是没有commit的内容

git diff HEAD //上面两个内容的合并

git reset --hard HEAD //撤销本地修改

echo $HOME //查看git config的HOME路径

export $HOME=/c/gitconfig //配置git config的HOME路径




Git 完整命令手册地址：http://git-scm.com/docs

PDF 版命令手册：github-git-cheat-sheet.pdf

