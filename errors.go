package errors

import "fmt"

var separator = ": "

// New returns an error with the supplied message.
// errCode optional parameter
func New(text string, errCode ...uint32) error {
	var code *uint32

	if len(errCode) != 0 {
		code = &errCode[0]
	}

	return &msg{
		text,
		code,
	}
}

// Errorf formats according to a format specifier and returns the string as a value that satisfies error
func Errorf(format string, args ...interface{}) error {
	return &msg{
		text: fmt.Sprintf(format, args...),
	}
}

// Wrap returns a new error that adds context to the original error
func Wrap(err error, text string) error {
	if err == nil {
		return nil
	}

	return &bucket{
		text,
		nil,
		addCause(err),
	}
}

// Wrap returns a new error that adds context to the original error with format specifier
func Wrapf(err error, format string, args ...interface{}) error {
	if err == nil {
		return nil
	}

	return wrapf(err, nil, format, args...)
}

// Cause will recursively retrieve the topmost error which does not implement `causer`, which is assumed to be the original cause
func Cause(err error) error {
	for err != nil {
		if cause, ok := err.(causer); ok && cause.Cause() != nil {
			err = cause.Cause()
			continue
		}

		break
	}

	return err
}

// SetSeparator replaces the default delimiter ": " with a custom one
// Warning, this operation is not safe for concurrent use, use in init function
func SetSeparator(sep string) {
	separator = sep
}

func wrapf(err error, code *uint32, format string, args ...interface{}) error {
	return &bucket{
		fmt.Sprintf(format, args...),
		code,
		addCause(err),
	}
}
