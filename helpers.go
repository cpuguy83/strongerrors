package errclass

import "context"

type errNotFound struct{ error }

func (errNotFound) NotFound() {}

// NotFound is a helper to create an error of the class with the same name from any error type
func NotFound(err error) error {
	if err == nil {
		return nil
	}
	return errNotFound{err}
}

type errInvalidArg struct{ error }

func (errInvalidArg) InvalidArgument() {}

// InvalidArgument is a helper to create an error of the class with the same name from any error type
func InvalidArgument(err error) error {
	if err == nil {
		return nil
	}
	return errInvalidArg{err}
}

type errConflict struct{ error }

func (errConflict) Conflict() {}

// Conflict is a helper to create an error of the class with the same name from any error type
func Conflict(err error) error {
	if err == nil {
		return nil
	}
	return errConflict{err}
}

type errUnauthorized struct{ error }

func (errUnauthorized) Unauthorized() {}

// Unauthorized is a helper to create an error of the class with the same name from any error type
func Unauthorized(err error) error {
	if err == nil {
		return nil
	}
	return errUnauthorized{err}
}

type errUnauthenticated struct{ error }

func (errUnauthorized) Unauthenticated() {}

// Unauthenticated is a helper to create an error of the class with the same name from any error type
func Unauthenticated(err error) error {
	if err == nil {
		return nil
	}
	return errUnauthenticated{err}
}

type errUnavailable struct{ error }

func (errUnavailable) Unavailable() {}

// Unavailable is a helper to create an error of the class with the same name from any error type
func Unavailable(err error) error {
	return errUnavailable{err}
}

type errForbidden struct{ error }

func (errForbidden) Forbidden() {}

// Forbidden is a helper to create an error of the class with the same name from any error type
func Forbidden(err error) error {
	if err == nil {
		return nil
	}
	return errForbidden{err}
}

type errSystem struct{ error }

func (errSystem) System() {}

// System is a helper to create an error of the class with the same name from any error type
func System(err error) error {
	if err == nil {
		return nil
	}
	return errSystem{err}
}

type errNotModified struct{ error }

func (errNotModified) NotModified() {}

// NotModified is a helper to create an error of the class with the same name from any error type
func NotModified(err error) error {
	if err == nil {
		return nil
	}
	return errNotModified{err}
}

type errAlreadyExists struct{ error }

func (errAlreadyExists) AlreadyExists() {}

// AlreadyExists is a helper to create an error of the class with the same name from any error type
func AlreadyExists(err error) error {
	if err == nil {
		return nil
	}
	return errAlreadyExists{err}
}

type errNotImplemented struct{ error }

func (errNotImplemented) NotImplemented() {}

// NotImplemented is a helper to create an error of the class with the same name from any error type
func NotImplemented(err error) error {
	if err == nil {
		return nil
	}
	return errNotImplemented{err}
}

type errUnknown struct{ error }

func (errUnknown) Unknown() {}

// Unknown is a helper to create an error of the class with the same name from any error type
func Unknown(err error) error {
	if err == nil {
		return nil
	}
	return errUnknown{err}
}

type errCancelled struct{ error }

func (errCancelled) Cancelled() {}

// Cancelled is a helper to create an error of the class with the same name from any error type
func Cancelled(err error) error {
	if err == nil {
		return nil
	}
	return errCancelled{err}
}

type errDeadline struct{ error }

func (errDeadline) DeadlineExceeded() {}

// Deadline is a helper to create an error of the class with the same name from any error type
func Deadline(err error) error {
	if err == nil {
		return nil
	}
	return errDeadline{err}
}

type errExhausted struct{ error }

func (errExhausted) Exhausted() {}

// Exhausted is a helper to create an error of the class with the same name from any error type
func Exhausted(err error) error {
	if err == nil {
		return nil
	}
	return errExhausted{err}
}

type errDataLoss struct{ error }

func (errDataLoss) DataLos() {}

// DataLoss is a helper to create an error of the class with the same name from any error type
func DataLoss(err error) error {
	if err == nil {
		return nil
	}
	return errDataLoss{err}
}

// FromContext returns the error class from the passed in context
func FromContext(ctx context.Context) error {
	e := ctx.Err()
	if e == nil {
		return nil
	}

	if e == context.Canceled {
		return Cancelled(e)
	}
	if e == context.DeadlineExceeded {
		return Deadline(e)
	}
	return Unknown(e)
}
