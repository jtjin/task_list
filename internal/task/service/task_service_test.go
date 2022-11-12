package service

import (
	"task_list/internal/task"
	mockTask "task_list/mock/task"
	"task_list/models"
	"task_list/models/apireq"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestTaskServiceCreate(t *testing.T) {
	// Arrange
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	mockTaskRepo := mockTask.NewMockRepository(ctl)
	taskService := NewTaskService(mockTaskRepo)
	name := "task 1"
	status := task.StatusIncomplete
	req := apireq.CreateTask{
		Name:   name,
		Status: status,
	}
	mockTask := &models.Task{
		Name:   name,
		Status: status,
	}
	mockTaskRepo.EXPECT().Insert(mockTask).Return(nil)

	// Act
	res, err := taskService.CreateTask(req)
	assert.Nil(t, err)
	assert.Equal(t, name, res.Result.Name)
	assert.Equal(t, status, res.Result.Status)
}

func TestTaskServiceListTask(t *testing.T) {
	// Arrange
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	mockTaskRepo := mockTask.NewMockRepository(ctl)
	taskService := NewTaskService(mockTaskRepo)
	mockTasks := []models.Task{
		{
			Id:     1,
			Name:   "task 1",
			Status: task.StatusIncomplete,
		},
		{
			Id:     2,
			Name:   "task 2",
			Status: task.StatusComplete,
		},
	}
	mockTaskRepo.EXPECT().FindAll().Return(mockTasks, nil)

	// Act
	res, err := taskService.ListTask()

	// Assert
	assert.Nil(t, err)
	assert.Len(t, res.Result, len(mockTasks))
	for i := range res.Result {
		assert.Equal(t, mockTasks[i].Id, res.Result[i].Id)
		assert.Equal(t, mockTasks[i].Name, res.Result[i].Name)
		assert.Equal(t, mockTasks[i].Status, res.Result[i].Status)
	}
}

func TestTaskServiceUpdate(t *testing.T) {
	// Arrange
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	mockTaskRepo := mockTask.NewMockRepository(ctl)
	taskService := NewTaskService(mockTaskRepo)

	id := 1
	nameBefore := "task before"
	nameAfter := "task after"
	statusBefore := task.StatusIncomplete
	statusAfter := task.StatusComplete

	req := apireq.UpdateTask{
		Id:     id,
		Name:   nameAfter,
		Status: &statusAfter,
	}

	// case 1 : id exists
	mockFindOneTask := &models.Task{
		Id:     id,
		Name:   nameBefore,
		Status: statusBefore,
	}
	mockTaskRepo.EXPECT().FindOneById(id).Return(mockFindOneTask, nil)
	mockUpdateTask := &models.Task{
		Id:     id,
		Name:   nameAfter,
		Status: statusAfter,
	}
	mockTaskRepo.EXPECT().Update(mockUpdateTask).Return(nil)

	// Act
	res, err := taskService.UpdateTask(req)
	assert.Nil(t, err)
	assert.Equal(t, nameAfter, res.Result.Name)
	assert.Equal(t, statusAfter, res.Result.Status)

	// case 2 : id not exists
	mockTaskRepo.EXPECT().FindOneById(id).Return(&models.Task{}, gorm.ErrRecordNotFound)

	// Act
	res, err = taskService.UpdateTask(req)
	assert.NotNil(t, err)
	assert.Equal(t, "task not exist", err.Error())
	assert.Nil(t, res)
}

func TestTaskServiceDelete(t *testing.T) {
	// Arrange
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	mockTaskRepo := mockTask.NewMockRepository(ctl)
	taskService := NewTaskService(mockTaskRepo)

	id := 1
	name := "task"
	status := task.StatusComplete

	req := apireq.DeleteTask{
		Id: id,
	}

	// case 1 : id exists
	mockTask := &models.Task{
		Id:     id,
		Name:   name,
		Status: status,
	}
	mockTaskRepo.EXPECT().FindOneById(id).Return(mockTask, nil)
	mockTaskRepo.EXPECT().Delete(mockTask).Return(nil)

	// Act
	err := taskService.DeleteTask(req)
	assert.Nil(t, err)

	// case 2 : id not exists
	mockTaskRepo.EXPECT().FindOneById(id).Return(&models.Task{}, gorm.ErrRecordNotFound)

	// Act
	err = taskService.DeleteTask(req)
	assert.NotNil(t, err)
	assert.Equal(t, "task not exist", err.Error())
}
