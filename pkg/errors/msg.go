package errors

var msg = map[int]string{
	ErrorParamInvalid:     "Wrong parameter format or invalid",
	ResourceNotFoundError: "Resource not found",

	DBQueryError:  "Database query error",
	DBInsertError: "Database insert error",
	DBUpdateError: "Database update error",
	DBDeleteError: "Database delete error",
}
