package middleware

import (
	"task_list/pkg/errors"

	"github.com/gin-gonic/gin"
)

func ErrorResponse() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		err := c.Errors.Last()

		if err == nil {
			return
		}

		switch err.Err.(type) {
		case *errors.AppError:
			appErr := err.Err.(*errors.AppError)
			c.AbortWithStatusJSON(appErr.GetStatus(), appErr.GetMsg())
			return
		default:
			c.AbortWithStatusJSON(400, err)
			return
		}
	}
}
