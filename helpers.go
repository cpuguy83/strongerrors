package errclass

type errNotFound struct{ error }

func (errNotFound) NotFound() {}

// NotFound is a helper to create an error of the class with the same name from any error type
func NotFound(err error) error {
	return errNotFound{err}
}

type errInvalidArg struct{ error }

func (errInvalidArg) InvalidArgument() {}

// InvalidArgument is a helper to create an error of the class with the same name from any error type
func InvalidArgument(err error) error {
	return errInvalidArg{err}
}

type errConflict struct{ error }

func (errConflict) Conflict() {}

// Conflict is a helper to create an error of the class with the same name from any error type
func Conflict(err error) error {
	return errConflict{err}
}

type errUnauthorized struct{ error }

func (errUnauthorized) Unauthorized() {}

// Unauthorized is a helper to create an error of the class with the same name from any error type
func Unauthorized(err error) error {
	return errUnauthorized{err}
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
	return errForbidden{err}
}

type errSystem struct{ error }

func (errSystem) System() {}

// System is a helper to create an error of the class with the same name from any error type
func System(err error) error {
	return errSystem{err}
}

type errNotModified struct{ error }

func (errNotModified) NotModified() {}

// NotModified is a helper to create an error of the class with the same name from any error type
func NotModified(err error) error {
	return errNotModified{err}
}

type errNotImplemented struct{ error }

func (errNotImplemented) NotImplemented() {}

// NotImplemented is a helper to create an error of the class with the same name from any error type
func NotImplemented(err error) error {
	return errNotImplemented{err}
}

type errUnknown struct{ error }

func (errUnknown) Unknown() {}

// Unknown is a helper to create an error of the class with the same name from any error type
func Unknown(err error) error {
	return errUnknown{err}
}
