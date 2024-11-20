package input

import (
	"bufio"
	"bytes"
	"strings"
	"testing"
)

func TestReadInput(t *testing.T) {
	reader := &Reader{scanner: bufio.NewReader(strings.NewReader("google.com\n"))}
	input, err := reader.ReadInput()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := "google.com"
	if input != expected {
		t.Fatalf("expected %s, got %s", expected, input)
	}
}

func TestReadInputEmpty(t *testing.T) {
	reader := &Reader{scanner: bufio.NewReader(strings.NewReader("\n"))}
	_, err := reader.ReadInput()
	if err == nil {
		t.Fatal("expected error, got none")
	}
}

func TestReadInputFailure(t *testing.T) {
	reader := &Reader{scanner: bufio.NewReader(bytes.NewReader([]byte{}))}
	_, err := reader.ReadInput()
	if err == nil {
		t.Fatal("expected error, got none")
	}
}
