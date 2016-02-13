package errors

import (
	"errors"
	"testing"
)

func TestError(t *testing.T) {
	var err error
	ioError := NewType("IOError")
	fileNotFoundError := SubType("FileNotFoundError", ioError)

	err = ioError.Errorf("No such file: %v", "test.txt")
	if !ioError.IsTypeOf(err) {
		t.Error("ioError should be IOError type.")
	}
	err = errors.New("some error")
	if ioError.IsTypeOf(err) {
		t.Error("ioError should not be IOError type.")
	}
	if ioError.IsTypeOf(nil) {
		t.Error("ioError should not be nil.")
	}

	err = fileNotFoundError.Errorf("No such file: %v", "test.txt")
	if !fileNotFoundError.IsTypeOf(err) {
		t.Error("fileNotFoundError should be FileNotFoundError type.")
	}
	if !ioError.IsTypeOf(err) {
		t.Error("fileNotFoundError should be IOError type.")
	}
}
