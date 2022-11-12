package errors

import "strconv"

type AppError struct {
	StatusCode int
	Code       string
	Msg        string
	CauseErr   error
}

type AppErrorMsg struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (e AppError) Error() string {
	return e.Msg
}

func NewAppErr(statusCode, code int, customMsg string, privateErr error) *AppError {
	er := &AppError{}
	s := strconv.Itoa(code)

	er.StatusCode = statusCode
	er.Code = s

	if customMsg != "" {
		er.Msg = customMsg
	} else {
		er.Msg = msg[code]
	}

	er.CauseErr = privateErr

	return er
}

func (e *AppError) GetStatus() int {
	return e.StatusCode
}

func (e *AppError) GetCode() int {
	code, _ := strconv.Atoi(e.Code)
	return code
}

func (e *AppError) GetMsg() *AppErrorMsg {
	return &AppErrorMsg{Message: e.Msg, Code: e.Code}
}
