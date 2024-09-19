package exceptions

import "fmt"


type ApplicationException struct {
	Title   string
	Message string
}

func (e *ApplicationException) Error() string {
	return fmt.Sprintf("%s: %s", e.Title, e.Message)
}


func NewApplicationException(title, message string) *ApplicationException {
	return &ApplicationException{
		Title:   title,
		Message: message,
	}
}
