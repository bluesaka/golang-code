// +build wireinject

package main

import (
	"github.com/google/wire"
	"go-code/study/di/wire/internal/cache"
	"go-code/study/di/wire/internal/config"
)

func InitApp() (*App, error) {
	// 调用wire.Build方法传入所有的依赖对象以及构建的最终目标对象
	panic(wire.Build(config.Provider, cache.Provider, NewApp))
}
