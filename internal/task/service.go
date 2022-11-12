package task

import (
	"task_list/models/apireq"
	"task_list/models/apires"
)

type Service interface {
	CreateTask(req apireq.CreateTask) (res *apires.CreateTask, err error)
	ListTask() (res *apires.ListTask, err error)
	UpdateTask(req apireq.UpdateTask) (res *apires.UpdateTask, err error)
	DeleteTask(req apireq.DeleteTask) (err error)
}
