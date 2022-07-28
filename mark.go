package errors

const unknown = "unknown"

type marker interface {
	mark(code *uint32)
	getMark() (uint32, bool)
}

// Mark add code to error
func Mark(err error, code uint32) error {
	if err == nil {
		return nil
	}

	if e, ok := err.(marker); ok {
		e.mark(&code)
		return err
	}

	return &bucket{
		"mark",
		&code,
		addCause(err),
	}
}

// Markf add code to error and adds context to the original error with format specifier
func Markf(err error, code uint32, format string, args ...interface{}) error {
	if err == nil {
		return nil
	}

	return wrapf(err, &code, format, args...)
}

// GetMark returns error code and http status code
func GetMark(err interface{}) (uint32, bool) {
	if e, ok := err.(marker); ok {
		return e.getMark()
	}

	return 0, false
}

func getMarkData(code uint32, httpStatuses map[uint32]map[string]interface{}) (string, int) {
	return getErrText(code, httpStatuses), getHttpStatus(code, httpStatuses)
}

func getHttpStatus(code uint32, httpStatuses map[uint32]map[string]interface{}) int {
	if _, ok := httpStatuses[code]; !ok {
		return 500
	}

	if _, ok := httpStatuses[code]["status"]; !ok {
		return 500
	}

	if status, ok := httpStatuses[code]["status"].(int); ok {
		return status
	}

	return 500
}

func getErrText(code uint32, httpStatuses map[uint32]map[string]interface{}) string {
	if _, ok := httpStatuses[code]; !ok {
		return unknown
	}

	if _, ok := httpStatuses[code]["text"]; !ok {
		return unknown
	}

	if status, ok := httpStatuses[code]["text"].(string); ok {
		return status
	}

	return unknown
}
