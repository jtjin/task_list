// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package api

import (
	"github.com/google/wire"
	"task_list/driver"
	"task_list/internal/task"
	"task_list/internal/task/repository"
	"task_list/internal/task/service"
)

// Injectors from wire.go:

func BuildTaskService() task.Service {
	db := driver.InitGorm()
	taskRepository := repository.NewMysqlTaskRepo(db)
	taskService := service.NewTaskService(taskRepository)
	return taskService
}

// wire.go:

var gormSet = wire.NewSet(driver.InitGorm)

var taskServiceSet = wire.NewSet(service.NewTaskService, repository.NewMysqlTaskRepo)