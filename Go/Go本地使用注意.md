## 1.win10下配置好环境变量

```
GOROOT:D:\Go

GOPATH:D:\gopath
```

## 2.go运行环境

go代码建立在`D:\gopath\src`，然后在里面创建不同的文件夹来安排自己的项目或者代码，同一个目录下不能存在多个`package`和`main`，不然会报错

## 3.VS code 使用

vscode是微软基于Electron和web技术构建的开源编辑器, 是一款很强大的编辑器。开源地址:https://github.com/Microsoft/vscode

1、安装Visual Studio Code 最新版

官方网站：https://code.visualstudio.com/ 下载Visual Studio Code 最新版，安装过程略。

2、安装Go插件

点击右边的 Extensions 图标 搜索Go插件 在插件列表中，选择 Go，进行安装，安装之后，系统会提示重启 Visual Studio Code。

建议把自动保存功能开启。开启方法为：选择菜单 File，点击 Auto save。

vscode代码设置可用于Go扩展。这些都可以在用户的喜好来设置或工作区设置（.vscode/settings.json）。

打开首选项-用户设置 settings.json:

```
{
    "files.autoSave"                     : "onFocusChange",
    "go.buildOnSave"                     : true,
    "go.lintOnSave"                      : true,
    "go.vetOnSave"                       : true,
    "go.buildTags"                       : "",
    "go.buildFlags"                      : [],
    "go.lintFlags"                       : [],
    "go.vetFlags"                        : [],
    "go.coverOnSave"                     : false,
    "go.useCodeSnippetsOnFunctionSuggest": false,
    "go.formatOnSave"                    : true,
    "go.formatTool"                      : "goreturns",
    "go.goroot"                          : "E:/Go",
    "go.gopath"                          : "E:/Gopath",
    "go.gocodeAutoBuild"                 : true
}
```
接着安装依赖包支持(网络不稳定,请直接到 GitHub [Golang](https://github.com/golang) 下载再移动到相关目录):

```
go get -u -v github.com/nsf/gocode
go get -u -v github.com/rogpeppe/godef
go get -u -v github.com/zmb3/gogetdoc
go get -u -v github.com/golang/lint/golint
go get -u -v github.com/lukehoban/go-outline
go get -u -v sourcegraph.com/sqs/goreturns
go get -u -v golang.org/x/tools/cmd/gorename
go get -u -v github.com/tpng/gopkgs
go get -u -v github.com/newhook/go-symbols
go get -u -v golang.org/x/tools/cmd/guru
go get -u -v github.com/cweill/gotests/...
```

打开首选项-工作区设置,配置launch.json:

```
{
    "version"       : "0.2.0",
    "configurations": [
        {
            "name"   : "Launch",
            "type"   : "go",
            "request": "launch",
            "mode"   : "debug",
            "program": "${fileDirname}",
            "port"   : 2345,
            "host"   : "127.0.0.1",
            "env"    : {},
            "args"   : [],
            "showLog": false
        }
    ]
}
```

