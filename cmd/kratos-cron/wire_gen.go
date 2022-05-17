// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-cron/internal/biz"
	"kratos-cron/internal/conf"
	"kratos-cron/internal/data"
	"kratos-cron/internal/server"
	"kratos-cron/internal/service"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, job *conf.Job, logger log.Logger) (*kratos.App, func(), error) {
	dataData, cleanup, err := data.NewData(confData, logger)
	if err != nil {
		return nil, nil, err
	}
	greeterRepo := data.NewGreeterRepo(dataData, logger)
	greeterUsecase := biz.NewGreeterUsecase(greeterRepo, logger)
	greeterService := service.NewGreeterService(greeterUsecase)
	httpServer := server.NewHTTPServer(confServer, greeterService, logger)
	grpcServer := server.NewGRPCServer(confServer, greeterService, logger)
	jobService := service.NewJobService(greeterUsecase)
	cronWorker := server.NewCronWorker(job, jobService)
	app := newApp(logger, httpServer, grpcServer, cronWorker)
	return app, func() {
		cleanup()
	}, nil
}