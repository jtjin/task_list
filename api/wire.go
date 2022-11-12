//go:build wireinject
// +build wireinject

package api

import (
	"task_list/driver"
	"task_list/internal/task"
	taskRepo "task_list/internal/task/repository"
	taskService "task_list/internal/task/service"

	"github.com/google/wire"
)

var gormSet = wire.NewSet(driver.InitGorm)

var taskServiceSet = wire.NewSet(taskService.NewTaskService, taskRepo.NewMysqlTaskRepo)

func BuildTaskService() task.Service {
	wire.Build(taskServiceSet, gormSet)
	return &taskService.TaskService{}
}
