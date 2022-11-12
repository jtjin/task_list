package seeds

import (
	"task_list/internal/task"
	"task_list/models"

	"gorm.io/gorm"
)

func CreateTask(orm *gorm.DB, task models.Task) error {
	return orm.Create(&task).Error
}

func AllTask() []Seed {
	return []Seed{
		{
			Name: "Create Tasks - name = task 1 / status = incomplete",
			Run: func(orm *gorm.DB) error {
				return CreateTask(orm, models.Task{
					Name:   "task 1",
					Status: task.StatusIncomplete,
				})
			},
		},
		{
			Name: "Create Tasks -  name = task  / status = complete",
			Run: func(orm *gorm.DB) error {
				return CreateTask(orm, models.Task{
					Name:   "task 2",
					Status: task.StatusComplete,
				})
			},
		},
	}
}
