package errors

type InternalServerError struct {
	Err error
}

func (e *InternalServerError) Error() string {
	return e.Err.Error()
}

type NotFoundError struct {
	Message string
}

func (e *NotFoundError) Error() string {
	return e.Message
}

type BadRequest struct {
	Message string
}

func (e *BadRequest) Error() string {
	return e.Message
}

type HTTPError struct {
	Err  error
	Code int
}
