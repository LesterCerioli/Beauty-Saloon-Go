package exceptions


type BadRequestException struct {
	ApplicationException
}

func (e *BadRequestException) Error() string {
	return e.ApplicationException.Error()
}


func NewBadRequestException(message string) *BadRequestException {
	return &BadRequestException{
		ApplicationException: *NewApplicationException("Bad Request", message),
	}
}
