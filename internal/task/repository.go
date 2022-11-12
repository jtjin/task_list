package task

import "task_list/models"

type Repository interface {
	Insert(task *models.Task) (err error)
	FindAll() (tasks []models.Task, err error)
	FindOneById(id int) (task *models.Task, err error)
	Update(task *models.Task) (err error)
	Delete(task *models.Task) (err error)
}
