package web

type ValidationError struct {
	Message string
}

func (e *ValidationError) Error() string {
	return e.Message
}

type InternalServerError struct {
	Message string
}

func (e *InternalServerError) Error() string {
	return e.Message
}
