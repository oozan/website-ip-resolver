package input

import (
	"bufio"
	"os"
	"strings"
	"website-ip-resolver/internal/handler"
)

type Reader struct {
	scanner *bufio.Reader
}

func NewReader() *Reader {
	return &Reader{scanner: bufio.NewReader(os.Stdin)}
}

func (r *Reader) ReadInput() (string, error) {
	input, err := r.scanner.ReadString('\n')
	if err != nil {
		return "", &errorhandler.CustomError{
			Message: "Failed to read input",
			Code:    500,
		}
	}
	input = strings.TrimSpace(input)
	if input == "" {
		return "", &errorhandler.CustomError{
			Message: "Input cannot be empty",
			Code:    400,
		}
	}
	return input, nil
}
