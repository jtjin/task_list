package api

import (
	"task_list/models/apireq"
	"task_list/pkg/errors"

	"github.com/gin-gonic/gin"
)

// @Summary List task
// @Tags Task
// @Accept json
// @Produce json
// @Success 200 {object} apires.ListTask
// @Failure 500 {object} errors.AppErrorMsg "{"code":"500001","message":"Database query error"}"
// @Router /tasks [get]
func ListTask(c *gin.Context) {
	res, err := BuildTaskService().ListTask()
	if err != nil {
		_ = c.Error(err)
		return
	}
	c.JSON(200, res)
}

// @Summary Create task
// @Tags Task
// @Accept json
// @Produce json
// @Param body body apireq.CreateTask true "Request 新增 Task (status 預設為 0)"
// @Success 201 {object} apires.CreateTask
// @Failure 400 {object} errors.AppErrorMsg "{"code":"400400","message":"Wrong parameter format or invalid"}"
// @Failure 500 {object} errors.AppErrorMsg "{"code":"500002","message":"Database insert error"}"
// @Router /task [post]
func CreateTask(c *gin.Context) {
	req := apireq.CreateTask{}
	if err := c.Bind(&req); err != nil {
		_ = c.Error(errors.NewAppErr(400, errors.ErrorParamInvalid, err.Error(), err))
		return
	}
	res, err := BuildTaskService().CreateTask(req)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(201, res)

}

// @Summary Update task
// @Tags Task
// @Accept json
// @Produce json
// @Param id path int true "task_id"
// @Param body body apireq.UpdateTask true "Request 修改 Task"
// @Success 200 {object} apires.UpdateTask
// @Failure 400 {object} errors.AppErrorMsg "{"code":"400400","message":"Wrong parameter format or invalid"}"
// @Failure 404 {object} errors.AppErrorMsg "{"code":"400404","message":"Resource not found"}"
// @Failure 500 {object} errors.AppErrorMsg "{"code":"500003","message":"Database update error"}"
// @Router /task/{id} [put]
func UpdateTask(c *gin.Context) {
	req := apireq.UpdateTask{}

	_ = c.ShouldBind(&req)
	if err := c.BindUri(&req); err != nil {
		_ = c.Error(errors.NewAppErr(400, errors.ErrorParamInvalid, err.Error(), err))
		return
	}
	res, err := BuildTaskService().UpdateTask(req)
	if err != nil {
		_ = c.Error(err)
		return
	}
	c.JSON(200, res)
}

// @Summary Delete task
// @Tags Task
// @Accept json
// @Produce json
// @Param id path int true "task_id"
// @Success 200 {string} string "{}"
// @Failure 400 {object} errors.AppErrorMsg "{"code":"400400","message":"Wrong parameter format or invalid"}"
// @Failure 404 {object} errors.AppErrorMsg "{"code":"400404","message":"Resource not found"}"
// @Failure 500 {object} errors.AppErrorMsg "{"code":"500004","message":"Database delete error"}"
// @Router /task/{id} [delete]
func DeleteTask(c *gin.Context) {
	req := apireq.DeleteTask{}

	if err := c.BindUri(&req); err != nil {
		_ = c.Error(errors.NewAppErr(400, errors.ErrorParamInvalid, err.Error(), err))
		return
	}
	err := BuildTaskService().DeleteTask(req)
	if err != nil {
		_ = c.Error(err)
		return
	}
	c.JSON(200, nil)
}
