package errorhandler

import (
	"errors"
	"testing"
)

func TestCustomError(t *testing.T) {
	err := &CustomError{
		Message: "Test error",
		Code:    404,
	}
	expected := "Error [404]: Test error"
	if err.Error() != expected {
		t.Fatalf("expected %s, got %s", expected, err.Error())
	}
}

func TestHandle(t *testing.T) {
	customErr := &CustomError{
		Message: "Handled error",
		Code:    400,
	}
	Handle(customErr)
	Handle(errors.New("Generic error"))
}
