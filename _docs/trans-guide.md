---
layout: default
title: 翻译指南
---

# Go 项目翻译指南
本文档描述了参与翻译的方法及注意事项。若需补充，请向
[本站代码库](https://github.com/golang-zh/golang-zh.github.io)
添加 issue，在进行过讨论后请在你的 fork 中提交修改并发送 PR（即 Pull Request，下同）。

## 前提需求
我们的翻译项目托管在 [Github](https://github.com/) 上，因此首先你需要有一个
Github 账户和一把足够长的梯子。为方便交流，你还需要一个电子邮件账户，我们推荐 Gmail。

Go 项目源码采用 [Git](http://git-scm.com/) 进行版本控制，因此你还需要安装 Git
并熟悉它的基本使用。我们推荐 [ProGit](http://git-scm.com/book/zh)
作为学习和参考的资料，当然[官方文档](http://git-scm.com/doc)也是个不错的选择。

## 申请参与
要想参与翻译很简单，直接在 [项目主页](https://github.com/golang-zh)
的列表中选择你想要参与的子项目，点击右上的 Fork 即可。在 fork
时，你可以选择一个名字以便和官方源码库区分（例如`Go-zh`）。在你第一次发送
PR 时，我们会把你添加到 Golang-zh 的 Translator 小组中，这样你就正式成为项目组成员了。

## 环境配置
在开始翻译之前，我们需要做一些配置，以方便后续的翻译工作。

### Git 配置
我们推荐你通过 SSH 访问 Github，这样安全性更好，而且省去了每次 push
输入密码的麻烦。配置方式详见 Github 的
[官方文档](https://help.github.com/articles/generating-ssh-keys)。

首先，你需要设置个人信息，只需在 Shell 中执行以下命令即可：

    git config --global user.name "你的用户名"
    git config --global user.email "你的邮件地址"

一般来说，Git 只需要基本的用户名和邮件地址就能正常工作。不过为了方便，你可以在
`~/.gitconfig` 中加入以下配置：

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
首先我们需要将自己 fork 的项目 clone 到本地。假设你已经配置好了 SSH，
请打开你的项目，点击右侧的 SSH 连接并复制你的 SSH clone URL（这里以 `OlingCat/Go-zh`
的 `git@github.com:OlingCat/Go-zh.git` 为例，你的会是另一个），在 Shell
中进入到想要放源码库的位置，接着执行：

    git clone git@github.com:OlingCat/Go-zh.git # 这就是前面复制的地址

稍等片刻后，当前目录下会出现一个 `Go-zh` 目录，这样我们就有了本地的代码库。
代码库默认分支为 `zh-master`，我们将在此分支下进行翻译。你可以执行 `cd Go-zh`
进入到该目录内再执行 `git st` 来确认，输出中应当会有 `On branch zh-master` 的字样。

在翻译过程中，我们还需要经常和主代码库同步，因此你还要设置远程代码库：

    git remote add upstream git@github.com:golang-zh/go.git

这里的 `upstream` 可以取任何名字（比如 `golang-zh`），只要你看到它能明白是什么就行。

为了让提交的历史记录看起来更简洁，我们需要为代码库设置自动 rebase：

    git config branch.autosetuprebase always

记得在每次 clone 代码库后都要执行。如果你觉得麻烦，而且确认不会影响其它代码库，
可以设为全局选项：

    git config --global branch.autosetuprebase always

## 准备翻译
翻译前需要认领翻译任务，这样既可以让大家知道你在翻译的文档，也可以避免重复翻译。

### 认领翻译任务
首先，你需要 fork [本站的源码](https://github.com/golang-zh/golang-zh.github.io)。
将自己的 fork clone 到本地后，添加用于同步的远程代码库：

    git remote add upstream git@github.com:golang-zh/golang-zh.github.io.git

接着在`_docs`目录下找到想要翻译的类别（一般是带`status-`前缀的`.md`文件），编辑它。
在文档内找到想要翻译的文档条目。若无人认领的话，直接在文档条目后添加认领信息即可。具体格式如下：

```
文档  // 译者: 状态[ (注释)]

其中：
文档 = 已列出的包名或文档名
译者 = 译者的ID[ & 同译者的ID]
状态 = *待译*|*翻译*|*校对*|*整理*|完成|…
注释 = 注释者ID: 具体说明

|  表示或者
[] 表示可选
```

若要认领任务，请注明译者并将状态标为`*待译*`；若已经开始翻译，请将状态改为`*翻译*`；
状态可以是上面列出状态以外的其它状态；非`完成`状态请使用`*星号突出*`格式标注。
译者可有多人，共同翻译者请在ID之间加`&`号。

例如，OlingCat 和 minux 共同翻译了 runtime 包的文档，现在正在校对。那么可以在 `status-pkg.md` 里这样写：

    ……
    runtime  // OlingCat & minux: *校对* (minux: extern.go, debug.go 仍需校对)
    ……

之后将修改的文件添加到 stage、提交、同步、push：

    git add _docs/status-pkg.md
    git ci -m "runtime // OlingCat & minux: *校对*"
    git pull upstream
    git push

最后发送 PR。等翻译完成后，只需改为：

    runtime  // OlingCat & minux: 完成

然后重复以上步骤即可。其它文档的状态更新步骤同上。

### 阅读翻译规范
翻译是个累活，需要注意的细节繁多。因此考虑到本文长度，我们不打算在这里叙述。
具体的方法及注意事项请阅读我们的[翻译规范](./trans-spec.html)。

现在一切准备就绪，我们可以开始进行翻译了。

## 翻译过程
在翻译过程中，难免会遇到拿不准的地方，这时请在 `Golang-zh` 对应的项目里创建 issue
发起讨论。你也可以在讨论中@某人的ID，这样可以让他在第一时间内看到并参与进来。
讨论结束后，请关闭该 issue。

### 初译
文档可以分段翻译（例如包文档可将文件作为单位），在一个阶段的翻译完成后，
需要将修改过的文件添加到 stage 并提交到本地代码库（这里以 `Go-zh` 里的
`src/runtime/extern.go` 为例）：

    git add src/runtime/extern.go
    git ci

此时 Git 会通过你设置的编辑器打开对应的提交描述文件，类似于这样：

```

# Please enter the commit message for your changes. Lines starting
# with '#' will be ignored, and an empty message aborts the commit.
# On branch zh-master
#
# Changes to be committed:
#   modified:   src/runtime/extern.go
#
```

该文件的第一行是空行，请在这行对你的更改进行简短的描述。若有必要，
请在换行后再插入一个空行，接着是详细的描述。如果有相关的引用，请在描述中提及。
我们约定的提交信息格式分为三部分：

```
包名或文档名: 简述。

以及一段
可选的详细描述。

# Please enter the commit message....
```

例如：

```
runtime: extern.go 校对完毕。

extern.go 文件内的部分文档仍需与官方同步更新，已标记 TODO；
debug.go 文件仍需校对。

# Please enter the commit message for your changes. Lines starting
# with '#' will be ignored, and an empty message aborts the commit.
# On branch zh-master
#
# Changes to be committed:
#   modified:   src/runtime/extern.go
#
```

之后保存关闭即可。当然，如果所有文件都需要提交，且无需详细描述，可以直接执行：

```
git ci -am "runtime: 校对完毕。"
```

之后，你就可以 push 并发送 PR 了。

### 审校
`Golang-zh` 项目组的成员会认真阅读并审校你的翻译。若有建议的更改，他们会直接在下方引用并回复。
请在本地修改文件后再次提交并发送 PR，审校者会再次对新的更改进行下一轮审校，如此重复。
当审校者们确认无误后，会回复 LGTM（Looks Good To Me：我很满意），这时管理组成员会接受此
PR 并合并到主代码库中，之后他人就可以同步你的更新了。

### 同步
