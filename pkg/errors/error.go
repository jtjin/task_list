package errors

import "strconv"

type AppError struct {
	StatusCode int
	Code       string
	Msg        string
	CauseErr   error
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
