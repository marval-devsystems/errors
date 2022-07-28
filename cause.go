package errors

type causer interface {
	Cause() error
}

type withCause struct {
	err error
}

func (wc *withCause) Cause() error {
	if wc == nil {
		return nil
	}

	return wc.err
}

func (wc *withCause) Error() string {
	return wc.err.Error()
}

func addCause(err error) *withCause {
	return &withCause{
		err,
	}
}
