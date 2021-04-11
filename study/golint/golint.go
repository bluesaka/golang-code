/**
一. Golint介绍
Golint 是一个源码检测工具用于检测代码规范
Golint 不同于gofmt, Gofmt用于代码格式化

Golint会对代码做以下几个方面检查：
- package注释 必须按照 “Package xxx 开头”
- package命名 不能有大写字母、下划线等特殊字符
- struct、interface等注释 必须按照指定格式开头
- struct、interface等命名
- 变量注释、命名
- 函数注释、命名
- 各种语法规范校验等

二. Golint安装
go get -u github.com/golang/lint/golint
```
go: found github.com/golang/lint/golint in github.com/golang/lint v0.0.0-20201208152925-83fdc39ff7b5
go get: github.com/golang/lint@v0.0.0-20201208152925-83fdc39ff7b5: parsing go.mod:
        module declares its path as: golang.org/x/lint
                but was required as: github.com/golang/lint
```
使用 go get -u golang.org/x/lint/golint 安装

ls $GOPATH/bin (可以发现已经有golint可执行文件)

三. Golint使用
golint检测代码有2种方式：
- golint file
- golint directory

golint main.go
golint cmd

四. Goland配置golint
添加tool：Goland -> Preferences -> Tools -> External Tools -> add golint (Program: $GOPATH/bin/golint, Argument: $FilePath$, Working Directory: $ProjectFileDir$)
设置快捷键：Goland -> Preferences -> Keymap -> External Tools -> External Tools -> Add Keyboard Shortcuts
*/

package golint

// 代码检查工具 golint golint.go
// comment on exported function Say should be of the form "Say ..."
// func parameter Http should be HTTP
func Say(Http int) int {
	return Http
}