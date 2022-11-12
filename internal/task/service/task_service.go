package service

import (
	"task_list/internal/task"
	"task_list/models"
	"task_list/models/apireq"
	"task_list/models/apires"
	"task_list/pkg/errors"

	"gorm.io/gorm"
)

type TaskService struct {
	taskRepo task.Repository
}

func NewTaskService(taskRepo task.Repository) task.Service {
	return &TaskService{
		taskRepo: taskRepo,
	}
}

func (s *TaskService) CreateTask(req apireq.CreateTask) (res *apires.Task, err error) {
	task := &models.Task{
		Name:   req.Name,
		Status: *req.Status,
	}
	err = s.taskRepo.Insert(task)
	res = &apires.Task{
		Id:     task.Id,
		Name:   task.Name,
		Status: task.Status,
	}
	if err != nil {
		return nil, errors.NewAppErr(500, errors.DBInsertError, "create task error.", err)
	}
	return res, nil
}

func (s *TaskService) ListTask() (res *apires.ListTask, err error) {
	tasks, err := s.taskRepo.FindAll()
	if err != nil {
		return nil, errors.NewAppErr(500, errors.DBQueryError, "find task error.", err)
	}
	res = new(apires.ListTask)
	for _, task := range tasks {
		res.Result = append(res.Result, apires.Task{
			Id:     task.Id,
			Name:   task.Name,
			Status: task.Status,
		})
	}
	return res, nil
}

func (s *TaskService) UpdateTask(req apireq.UpdateTask) (res *apires.Task, err error) {
	// check data exists
	task, err := s.taskRepo.FindOneById(req.Id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.NewAppErr(404, errors.ResourceNotFoundError, "task not exist", err)
		}
		return nil, errors.NewAppErr(500, errors.DBUpdateError, "find task error", err)
	}
	task.Name = req.Name
	task.Status = *req.Status
	err = s.taskRepo.Update(task)
	if err != nil {
		return nil, errors.NewAppErr(500, errors.DBUpdateError, "update task error", err)
	}
	res = &apires.Task{
		Id:     task.Id,
		Name:   task.Name,
		Status: task.Status,
	}
	return res, nil
}

func (s *TaskService) DeleteTask(req apireq.DeleteTask) (err error) {
	// check data exists
	task, err := s.taskRepo.FindOneById(req.Id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.NewAppErr(404, errors.ResourceNotFoundError, "task not exist", err)
		}
		return errors.NewAppErr(500, errors.DBUpdateError, "find task error", err)
	}
	err = s.taskRepo.Delete(task)
	if err != nil {
		return errors.NewAppErr(500, errors.DBDeleteError, "delete task error", err)
	}
	return nil
}
