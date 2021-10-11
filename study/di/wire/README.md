## Wire

> link: https://go-kratos.dev/docs/guide/wire
> link: https://go-kratos.dev/blog/go-project-wire

```
Google Wire 是一个灵活的依赖注入工具，通过自动生成代码的方式在编译期完成依赖注入。

在各个组件之间的依赖关系中，通常鼓励显式初始化，而不是全局变量传递。

所以通过 Wire 进行初始化代码，可以很好地解决组件之间的耦合，以及提高代码维护性。
```

### 安装

```
go get github.com/google/wire/cmd/wire  安装wire命令行工具
```

### 工作原理

> Wire 具有两个基本概念：Provider 和 Injector。

- Provider

> Provider 是一个普通的 Go Func ，负责创建对象的方法，例如NewRedis等

- Injector

> 根据对象依赖，构造目的对象的方法


### 示例

```
|--cmd
	|--main.go
	|--wire.go
|--config
	|--config.yaml
|--internal
	|--config
		|--config.go
	|--db
		|--db.go
```

### 生成wire_gen.go文件

```
在cmd目录下执行wire命令
```