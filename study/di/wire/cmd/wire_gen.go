// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"go-code/study/di/wire/internal/cache"
	"go-code/study/di/wire/internal/config"
)

// Injectors from wire.go:

func InitApp() (*App, error) {
	configConfig, err := config.NewConfig()
	if err != nil {
		return nil, err
	}
	pool := cache.NewRedis(configConfig)
	app := NewApp(pool)
	return app, nil
}