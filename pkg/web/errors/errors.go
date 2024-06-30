package errors

type InternalServerError struct {
	Err error
}

func (e *InternalServerError) Error() string {
	return e.Err.Error()
}
