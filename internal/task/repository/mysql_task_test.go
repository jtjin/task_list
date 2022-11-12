package repository

import (
	"task_list/driver"
	"task_list/internal/task"
	"task_list/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTaskRepoInsert(t *testing.T) {
	// Arrange
	orm := driver.InitGorm()
	taskRepo := NewMysqlTaskRepo(orm)
	insertTask := &models.Task{
		Name:   "task insert",
		Status: task.StatusIncomplete,
	}

	// Act
	err := taskRepo.Insert(insertTask)

	// Assert
	assert.Nil(t, err)
	assert.NotZero(t, insertTask.Id)

	// Teardown
	err = orm.Delete(insertTask).Error
	assert.Nil(t, err)
}

func TestTaskRepoFindAll(t *testing.T) {
	// Arrange
	orm := driver.InitGorm()
	taskRepo := NewMysqlTaskRepo(orm)

	// Act
	tasks, err := taskRepo.FindAll()

	// Assert
	assert.Nil(t, err)
	assert.Len(t, tasks, 2)
}

func TestTaskRepoFindOneById(t *testing.T) {
	// Arrange
	orm := driver.InitGorm()
	taskRepo := NewMysqlTaskRepo(orm)

	// case 1 : id exists
	// Act
	id := 1
	task, err := taskRepo.FindOneById(id)

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, id, task.Id)
	assert.NotEmpty(t, task.Name)

	// case 2 : id not exists
	// Act
	id = 999
	task, err = taskRepo.FindOneById(id)

	// Assert
	assert.NotNil(t, err)
	assert.Zero(t, task.Id)
	assert.Empty(t, task.Name)
}

func TestTaskRepoUpdate(t *testing.T) {
	// Arrange
	orm := driver.InitGorm()
	taskRepo := NewMysqlTaskRepo(orm)

	// 先新增一筆 task
	insertTask := &models.Task{
		Name:   "task insert",
		Status: task.StatusIncomplete,
	}
	err := orm.Create(insertTask).Error
	assert.Nil(t, err)

	condition := &models.Task{
		Id:     insertTask.Id,
		Name:   "task update",
		Status: task.StatusComplete,
	}
	// Act
	err = taskRepo.Update(condition)

	// Assert
	assert.Nil(t, err)
	taskAfterUpdate := new(models.Task)
	err = orm.Where("id = ?", insertTask.Id).First(taskAfterUpdate).Error
	assert.Nil(t, err)
	assert.Equal(t, condition.Id, taskAfterUpdate.Id)
	assert.Equal(t, condition.Name, taskAfterUpdate.Name)
	assert.Equal(t, condition.Status, taskAfterUpdate.Status)

	// Teardown
	err = orm.Delete(insertTask).Error
	assert.Nil(t, err)
}

func TestTaskRepoDelete(t *testing.T) {
	// Arrange
	orm := driver.InitGorm()
	taskRepo := NewMysqlTaskRepo(orm)

	// 先新增一筆 task
	insertTask := &models.Task{
		Name:   "task insert",
		Status: task.StatusIncomplete,
	}
	err := orm.Create(insertTask).Error
	assert.Nil(t, err)

	condition := &models.Task{
		Id: insertTask.Id,
	}
	// Act
	err = taskRepo.Delete(condition)

	// Assert
	assert.Nil(t, err)
	taskAfterDelete := new(models.Task)
	err = orm.Where("id = ?", insertTask.Id).First(taskAfterDelete).Error
	assert.NotNil(t, err)
	assert.Zero(t, taskAfterDelete.Id)
	assert.Empty(t, taskAfterDelete.Name)
}
