package api_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"task_list/models"
	"task_list/models/apireq"
	"task_list/models/apires"
	"task_list/pkg/helper"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListTask(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, "/tasks", bytes.NewBuffer(nil))
	req.Header.Set("Content-Type", "application/json")
	response := httptest.NewRecorder()
	routes.ServeHTTP(response, req)

	assert.Equal(t, 200, response.Code)

	var result apires.ListTask
	bodyBytes := response.Body.Bytes()
	err := json.Unmarshal(bodyBytes, &result)

	assert.Nil(t, err)
	assert.Len(t, result.Result, 2)
}

func TestCreateTask(t *testing.T) {
	var (
		Success        = 201
		ValidationFail = 400

		StatusIncomplete = 0
		StatusComplete   = 1
		StatusUnknown    = 999
	)
	tests := []struct {
		name         string
		body         apireq.CreateTask
		responseCode int
		errorMsg     string
	}{
		{
			name: "create task: success - without status",
			body: apireq.CreateTask{
				Name: "task",
			},
			responseCode: Success,
		},
		{
			name: "create task: success - status = incomplete",
			body: apireq.CreateTask{
				Name:   "task",
				Status: StatusIncomplete,
			},
			responseCode: Success,
		},
		{
			name: "create task: success - status = complete",
			body: apireq.CreateTask{
				Name:   "task",
				Status: StatusComplete,
			},
			responseCode: Success,
		},
		{
			name: "create task: fail - status = unknown",
			body: apireq.CreateTask{
				Name:   "task",
				Status: StatusUnknown,
			},
			responseCode: ValidationFail,
			errorMsg:     `{"code":"400400","message":"Key: 'CreateTask.Status' Error:Field validation for 'Status' failed on the 'oneof' tag"}`,
		},
		{
			name: "create task: fail - without name",
			body: apireq.CreateTask{
				Status: StatusComplete,
			},
			responseCode: ValidationFail,
			errorMsg:     `{"code":"400400","message":"Key: 'CreateTask.Name' Error:Field validation for 'Name' failed on the 'required' tag"}`,
		},
		{
			name: "create task: fail - len(name) > 250",
			body: apireq.CreateTask{
				Name:   helper.RandString(251),
				Status: StatusComplete,
			},
			responseCode: ValidationFail,
			errorMsg:     `{"code":"400400","message":"Key: 'CreateTask.Name' Error:Field validation for 'Name' failed on the 'max' tag"}`,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			body, err := json.Marshal(&test.body)
			assert.Equal(t, nil, err)

			req, _ := http.NewRequest(http.MethodPost, "/task", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			response := httptest.NewRecorder()
			routes.ServeHTTP(response, req)

			assert.Equal(t, test.responseCode, response.Code)

			switch response.Code {
			case Success:
				var result apires.CreateTask
				bodyBytes := response.Body.Bytes()
				err = json.Unmarshal(bodyBytes, &result)
				assert.Nil(t, err)
				assert.NotZero(t, result.Result.Id)
				assert.Equal(t, test.body.Name, result.Result.Name)
				assert.Equal(t, test.body.Status, result.Result.Status)

				// Teardown
				err = orm.Where("id = ?", result.Result.Id).Delete(models.Task{}).Error
				assert.Nil(t, err)
			case ValidationFail:
				assert.Equal(t, test.errorMsg, response.Body.String())
			}
		})
	}
}

func TestUpdateTask(t *testing.T) {
	var (
		Success          = 200
		ValidationFail   = 400
		NotFound         = 404
		StatusIncomplete = 0
		StatusComplete   = 1
		StatusUnknown    = 999
	)
	tests := []struct {
		name         string
		body         apireq.UpdateTask
		responseCode int
		errorMsg     string
	}{
		{
			name: "update task: success - status = 0",
			body: apireq.UpdateTask{
				Id:     1,
				Name:   "task",
				Status: &StatusIncomplete,
			},
			responseCode: Success,
		},
		{
			name: "update task: success - status = 1",
			body: apireq.UpdateTask{
				Id:     1,
				Name:   "task",
				Status: &StatusComplete,
			},
			responseCode: Success,
		},
		{
			name: "update task: fail - without id",
			body: apireq.UpdateTask{
				Name:   "task",
				Status: &StatusComplete,
			},
			responseCode: ValidationFail,
			errorMsg:     `{"code":"400400","message":"Key: 'UpdateTask.Id' Error:Field validation for 'Id' failed on the 'required' tag"}`,
		},
		{
			name: "update task: fail - without status",
			body: apireq.UpdateTask{
				Id:   1,
				Name: "task",
			},
			responseCode: ValidationFail,
			errorMsg:     `{"code":"400400","message":"Key: 'UpdateTask.Status' Error:Field validation for 'Status' failed on the 'required' tag"}`,
		},
		{
			name: "update task: fail - without name",
			body: apireq.UpdateTask{
				Id:     1,
				Status: &StatusComplete,
			},
			responseCode: ValidationFail,
			errorMsg:     `{"code":"400400","message":"Key: 'UpdateTask.Name' Error:Field validation for 'Name' failed on the 'required' tag"}`,
		},
		{
			name: "update task: fail - status = unknown",
			body: apireq.UpdateTask{
				Id:     1,
				Name:   "task",
				Status: &StatusUnknown,
			},
			responseCode: ValidationFail,
			errorMsg:     `{"code":"400400","message":"Key: 'UpdateTask.Status' Error:Field validation for 'Status' failed on the 'oneof' tag"}`,
		},
		{
			name: "update task: fail - id not exists",
			body: apireq.UpdateTask{
				Id:     999,
				Name:   "task",
				Status: &StatusComplete,
			},
			responseCode: NotFound,
			errorMsg:     `{"code":"400404","message":"task not exist"}`,
		},
		{
			name: "update task: fail - len(name) > 250",
			body: apireq.UpdateTask{
				Id:     1,
				Name:   helper.RandString(251),
				Status: &StatusComplete,
			},
			responseCode: ValidationFail,
			errorMsg:     `{"code":"400400","message":"Key: 'UpdateTask.Name' Error:Field validation for 'Name' failed on the 'max' tag"}`,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			body, err := json.Marshal(&test.body)
			assert.Equal(t, nil, err)

			// 紀錄 task 的原始數據
			originTask := models.Task{Id: test.body.Id}
			_ = orm.First(&originTask).Error

			req, _ := http.NewRequest(http.MethodPut, fmt.Sprintf("/task/%d", test.body.Id), bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			response := httptest.NewRecorder()
			routes.ServeHTTP(response, req)

			assert.Equal(t, test.responseCode, response.Code)

			switch response.Code {
			case Success:
				var result apires.UpdateTask
				bodyBytes := response.Body.Bytes()
				err = json.Unmarshal(bodyBytes, &result)
				assert.Nil(t, err)
				assert.Equal(t, test.body.Id, result.Result.Id)
				assert.Equal(t, test.body.Name, result.Result.Name)
				assert.Equal(t, *test.body.Status, result.Result.Status)

				// Teardown
				err = orm.Select("*").Updates(&originTask).Error
				assert.Nil(t, err)
			case ValidationFail, NotFound:
				assert.Equal(t, test.errorMsg, response.Body.String())
			}
		})
	}
}

func TestDeleteTask(t *testing.T) {
	var (
		Success        = 200
		ValidationFail = 400
		NotFound       = 404
	)
	tests := []struct {
		name         string
		body         apireq.DeleteTask
		responseCode int
		errorMsg     string
	}{
		{
			name: "delete task: success",
			body: apireq.DeleteTask{
				Id: 1,
			},
			responseCode: Success,
		},
		{
			name:         "delete task: fail - without id",
			body:         apireq.DeleteTask{},
			responseCode: ValidationFail,
			errorMsg:     `{"code":"400400","message":"Key: 'DeleteTask.Id' Error:Field validation for 'Id' failed on the 'required' tag"}`,
		},
		{
			name: "delete task: fail - id not exists",
			body: apireq.DeleteTask{
				Id: 999,
			},
			responseCode: NotFound,
			errorMsg:     `{"code":"400404","message":"task not exist"}`,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			body, err := json.Marshal(&test.body)
			assert.Equal(t, nil, err)

			// 紀錄 task 的原始數據
			originTask := models.Task{Id: test.body.Id}
			_ = orm.First(&originTask).Error

			req, _ := http.NewRequest(http.MethodDelete, fmt.Sprintf("/task/%d", test.body.Id), bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			response := httptest.NewRecorder()
			routes.ServeHTTP(response, req)

			assert.Equal(t, test.responseCode, response.Code)

			switch response.Code {
			case Success:
				// Teardown
				err = orm.Create(&originTask).Error
				assert.Nil(t, err)
			case ValidationFail, NotFound:
				assert.Equal(t, test.errorMsg, response.Body.String())
			}
		})
	}
}
