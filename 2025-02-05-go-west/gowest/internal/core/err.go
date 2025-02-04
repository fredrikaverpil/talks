package core

type CustomError struct {
	Message string
}

func (e *CustomError) Error() string {
	return e.Message
}
