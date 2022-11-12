package route

import (
	"task_list/config"
	"task_list/docs"
	"task_list/driver"
	"task_list/middleware"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Init() *gin.Engine {
	// init config and gorm
	config.InitConfig()
	driver.InitGorm()

	// init gin server
	r := gin.Default()

	// init swagger
	docs.SwaggerInfo.BasePath = ""
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// init router
	r.Use(middleware.ErrorResponse())
	TaskRoute(r)

	return r
}
