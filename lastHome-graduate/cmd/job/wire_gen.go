// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"blog/internal/biz"
	"blog/internal/cache/redis"
	"blog/internal/conf"
	"blog/internal/data"
	"blog/internal/mq/kafka"
	"blog/internal/service"
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

// Injectors from wire.go:

// initApp init kratos application.
func initApp(confData *conf.Data, cache *conf.Cache, messageQueue *conf.MessageQueue, logger log.Logger) (func(context.Context) error, func(), error) {
	dataData, cleanup, err := data.NewData(confData, logger)
	if err != nil {
		return nil, nil, err
	}
	summaryRepo := data.NewSummaryRepo(dataData, logger)
	summaryUsecase := biz.NewSummaryUsecase(summaryRepo, logger)
	radixRC3, err := redis.NewRedisClient(cache)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	kafkaClient := kafka.NewKafkaClient(messageQueue)
	jobService := service.NewJobService(summaryUsecase, radixRC3, kafkaClient, logger)
	v := newApp(jobService)
	return v, func() {
		cleanup()
	}, nil
}
