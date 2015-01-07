---
layout: default
title: 翻译指南
---

# Go 项目翻译指南
本文档描述了参与翻译的方法及注意事项。若需补充，请向 https://github.com/golang-zh/golang-zh.github.io 发送 Pull Request。

## 前提需求
我们的翻译项目托管在 [Github](https://github.com/) 上，因此首先你需要有一个 Github 账户。为方便交流，你还需要一个电子邮件账户，我们推荐 Gmail。

Go 项目源码采用 [Git](http://git-scm.com/) 进行版本控制，因此你还需要安装 Git 并熟悉它的基本使用。我们推荐 [ProGit](http://git-scm.com/book/zh) 作为学习和参考的资料，当然[官方文档](http://git-scm.com/doc)也是个不错的选择。

## 申请参与
要想参与翻译很简单，直接在[项目主页](https://github.com/Golang-zh)的列表中选择你想要参与的子项目，点击右上的 Fork 即可。在 fork 时你可以选择一个名字以便和官方源码库区分（例如`Go-zh`）。在你第一次发送 Pull Request 时，我们会把你添加到 Golang-zh 的 Translator 小组中。

## 环境配置
在开始翻译之前，我们需要做一些配置，以方便后续的翻译工作。

### Git 配置
我们推荐你通过 SSH 访问 Github，这样安全性好，而且不用每次输入密码。配置方式详见 Github 的[官方文档](https://help.github.com/articles/generating-ssh-keys)。

首先，你需要设置个人信息，只需在 Shell 中执行以下命令即可：

    git config --global user.name "你的用户名"
    git config --global user.email "你的邮件地址"

一般来说，Git 只需要基本的用户名和邮件地址，就能正常工作。不过为了方便，你可以在 `~/.gitconfig` 中加入以下配置：

```
[user]
	name = 你的用户名
	email = 你的邮件地址
[core]
	editor = 你的惯用编辑器
	autocrlf = input
[color]
	ui = true
	status = auto
	branch = auto
[branch]
	autosetuprebase = always
[alias]
	co = checkout
	br = branch
	ci = commit
	st = status
	last = log -1 HEAD
	unstage = reset HEAD --
```

配置详情请参考官方文档。

### 本地代码库配置
首先我们需要将自己 fork 的项目 clone 到本地。假设你已经配置好了SSH，请打开你的项目，点击右侧的 SSH 连接并复制你的 SSH clone URL（这里以 `OlingCat/Go-zh` 的 `git@github.com:OlingCat/Go-zh.git` 为例，你的会是另一个），在 Shell 中进入到想要放源码库的位置，接着执行：

    # clone 远程代码库到本地
    git clone git@github.com:OlingCat/Go-zh.git

    # 进入本地代码库
    cd Go-zh

稍等片刻后，我们就有了本地的代码库。当前目录下会出现一个 `Go-zh` 目录，默认的分支为 `zh-master`，我们将在此分支下进行翻译。你可以在此目录内执行 `git st` 来确认，输出中应当会有 `On branch zh-master` 字样。

在翻译过程中，我们还需要经常和主代码库同步，因此要设置远程代码库：

    git remote add golang-zh git@github.com:golang-zh/go.git

这里的 `golang-zh` 可以取任何名字，只要看到就能明白它的用途即可。

现在，我们就可以开始进行翻译了。

## 开始翻译

### 认领翻译任务
