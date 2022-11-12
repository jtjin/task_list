package route

import (
	"task_list/api"

	"github.com/gin-gonic/gin"
)

func TaskRoute(r *gin.Engine) {
	r.GET("/tasks", api.ListTask)
	r.POST("/task", api.CreateTask)
	r.PUT("/task/:id", api.UpdateTask)
	r.DELETE("/task/:id", api.DeleteTask)
}
