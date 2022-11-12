package errors

const (
	ErrorParamInvalid     = 400400
	ResourceNotFoundError = 400404

	UnknownError  = 500000
	DBQueryError  = 500001
	DBInsertError = 500002
	DBUpdateError = 500003
	DBDeleteError = 500004
)
