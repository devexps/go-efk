// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/devexps/go-efk/api/gen/go/common/conf"
	"github.com/devexps/go-efk/helloworld/internal/biz"
	"github.com/devexps/go-efk/helloworld/internal/data"
	"github.com/devexps/go-efk/helloworld/internal/server"
	"github.com/devexps/go-efk/helloworld/internal/service"
	"github.com/devexps/go-micro/v2"
	"github.com/devexps/go-micro/v2/log"
	"github.com/devexps/go-micro/v2/registry"
)

// Injectors from wire.go:

// initApp init micro application.
func initApp(logger log.Logger, registrar registry.Registrar, bootstrap *conf.Bootstrap) (*micro.App, func(), error) {
	dataData, cleanup, err := data.NewData(bootstrap, logger)
	if err != nil {
		return nil, nil, err
	}
	greeterRepo := data.NewGreeterRepo(dataData, logger)
	greeterUseCase := biz.NewGreeterUseCase(greeterRepo, logger)
	greeterService := service.NewGreeterService(greeterUseCase)
	httpServer := server.NewHTTPServer(bootstrap, logger, greeterService)
	grpcServer := server.NewGRPCServer(bootstrap, logger, greeterService)
	app := newApp(logger, httpServer, grpcServer, registrar)
	return app, func() {
		cleanup()
	}, nil
}
