package errorhandler

import "fmt"

type CustomError struct {
	Message string
	Code    int
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("Error [%d]: %s", e.Code, e.Message)
}

func Handle(err error) {
	if customErr, ok := err.(*CustomError); ok {
		fmt.Printf("Handled: %s\n", customErr.Error())
	} else if err != nil {
		fmt.Printf("Unhandled Error: %s\n", err.Error())
	}
}
