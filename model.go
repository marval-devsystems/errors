package errors

type msg struct {
	text string
	code *uint32
}

type bucket struct {
	text string
	code *uint32
	*withCause
}

func (e *msg) Error() string {
	return e.text
}

func (e *bucket) Error() string {
	text := e.text
	cause := e.withCause

	for cause != nil {
		text = text + separator + cause.Error()

		if err, ok := cause.err.(*withCause); ok {
			cause = err
			continue
		}

		break
	}

	return text
}

func (e *msg) mark(code *uint32) {
	e.code = code
}

func (e *bucket) mark(code *uint32) {
	e.code = code
}

func (e *msg) getMark() (uint32, bool) {
	if e.code == nil {
		return 0, false
	}

	code := *e.code

	return code, true
}

func (e *bucket) getMark() (uint32, bool) {
	if e.code == nil {
		return 0, false
	}

	code := *e.code

	return code, true
}

func (e *bucket) Unwrap() error {
	return e.err
}
