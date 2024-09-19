package exceptions


type NotFoundException struct {
	ApplicationException
}

func (e *NotFoundException) Error() string {
	return e.ApplicationException.Error()
}


func NewNotFoundException(message string) *NotFoundException {
	return &NotFoundException{
		ApplicationException: *NewApplicationException("Not Found", message),
	}
}
