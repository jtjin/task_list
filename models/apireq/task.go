package apireq

type CreateTask struct {
	Name   string `json:"name" binding:"required"`
	Status int    `json:"status" binding:"oneof=0 1"`
}

type UpdateTask struct {
	Id     int    `uri:"id" binding:"required" swaggerignore:"true"`
	Name   string `json:"name" binding:"required"`
	Status *int   `json:"status" binding:"required,oneof=0 1"`
}

type DeleteTask struct {
	Id int `uri:"id" binding:"required" swaggerignore:"true"`
}
