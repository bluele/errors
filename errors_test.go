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
		t.Error("ioError should be io type.")
	}
	err = errors.New("some error")
	if ioError.IsTypeOf(err) {
		t.Error("ioError should not be io type.")
	}
	if ioError.IsTypeOf(nil) {
		t.Error("ioError should not be nil.")
	}

	err = fileNotFoundError.Errorf("No such file: %v", "test.txt")
	if !fileNotFoundError.IsTypeOf(err) {
		t.Error("fileNotFoundError should be fileNotFound type.")
	}
	if !ioError.IsTypeOf(err) {
		t.Error("fileNotFoundError should be inherited by fileNotFound type.")
	}
}
