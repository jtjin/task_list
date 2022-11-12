package repository

import (
	"task_list/internal/task"
	"task_list/models"

	"gorm.io/gorm"
)

type TasksRepo struct {
	orm *gorm.DB
}

func NewMysqlTaskRepo(orm *gorm.DB) task.Repository {
	return &TasksRepo{
		orm: orm,
	}
}

func (r *TasksRepo) Insert(task *models.Task) (err error) {
	return r.orm.Create(task).Error
}

func (r *TasksRepo) FindAll() (tasks []models.Task, err error) {
	err = r.orm.Find(&tasks).Error
	return tasks, err
}

func (r *TasksRepo) FindOneById(id int) (task *models.Task, err error) {
	task = new(models.Task)
	err = r.orm.Where("id = ?", id).First(task).Error
	return task, err
}

func (r *TasksRepo) Update(task *models.Task) (err error) {
	return r.orm.Select("*").Updates(task).Error
}

func (r *TasksRepo) Delete(task *models.Task) (err error) {
	return r.orm.Delete(task).Error
}
