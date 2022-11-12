package route

import (
	"task_list/config"
	"task_list/driver"

	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	// init config and gorm
	config.InitConfig()
	driver.InitGorm()

	// init gin server
	r := gin.Default()

	return r
}
