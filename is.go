package strongerrors

type causer interface {
	Cause() error
}

func getImplementer(err error) error {
	switch e := err.(type) {
	case
		ErrNotFound,
		ErrInvalidArgument,
		ErrConflict,
		ErrUnauthorized,
		ErrUnauthenticated,
		ErrUnavailable,
		ErrForbidden,
		ErrSystem,
		ErrNotModified,
		ErrAlreadyExists,
		ErrNotImplemented,
		ErrCancelled,
		ErrDeadline,
		ErrDataLoss,
		ErrExhausted,
		ErrUnknown:
		return err
	case causer:
		return getImplementer(e.Cause())
	default:
		return err
	}
}

// IsNotFound returns if the passed in error is an ErrNotFound
func IsNotFound(err error) bool {
	_, ok := getImplementer(err).(ErrNotFound)
	return ok
}

// IsInvalidArgument returns if the passed in error is an ErrInvalidParameter
func IsInvalidArgument(err error) bool {
	_, ok := getImplementer(err).(ErrInvalidArgument)
	return ok
}

// IsConflict returns if the passed in error is an ErrConflict
func IsConflict(err error) bool {
	_, ok := getImplementer(err).(ErrConflict)
	return ok
}

// IsUnauthorized returns if the the passed in error is an ErrUnauthorized
func IsUnauthorized(err error) bool {
	_, ok := getImplementer(err).(ErrUnauthorized)
	return ok
}

// IsUnauthenticated returns if the the passed in error is an ErrUnauthenticated
func IsUnauthenticated(err error) bool {
	_, ok := getImplementer(err).(ErrUnauthenticated)
	return ok
}

// IsUnavailable returns if the passed in error is an ErrUnavailable
func IsUnavailable(err error) bool {
	_, ok := getImplementer(err).(ErrUnavailable)
	return ok
}

// IsForbidden returns if the passed in error is an ErrForbidden
func IsForbidden(err error) bool {
	_, ok := getImplementer(err).(ErrForbidden)
	return ok
}

// IsSystem returns if the passed in error is an ErrSystem
func IsSystem(err error) bool {
	_, ok := getImplementer(err).(ErrSystem)
	return ok
}

// IsNotModified returns if the passed in error is a NotModified error
func IsNotModified(err error) bool {
	_, ok := getImplementer(err).(ErrNotModified)
	return ok
}

// IsAlreadyExists returns if the passed in error is a AlreadyExists error
func IsAlreadyExists(err error) bool {
	_, ok := getImplementer(err).(ErrAlreadyExists)
	return ok
}

// IsNotImplemented returns if the passed in error is an ErrNotImplemented
func IsNotImplemented(err error) bool {
	_, ok := getImplementer(err).(ErrNotImplemented)
	return ok
}

// IsUnknown returns if the passed in error is an ErrUnknown
func IsUnknown(err error) bool {
	_, ok := getImplementer(err).(ErrUnknown)
	return ok
}

// IsCancelled returns if the passed in error is an ErrCancelled
func IsCancelled(err error) bool {
	_, ok := getImplementer(err).(ErrCancelled)
	return ok
}

// IsDeadline returns if the passed in error is an ErrDeadline
func IsDeadline(err error) bool {
	_, ok := getImplementer(err).(ErrDeadline)
	return ok
}

// IsExhausted returns if the passed in error is an ErrDeadline
func IsExhausted(err error) bool {
	_, ok := getImplementer(err).(ErrExhausted)
	return ok
}

// IsDataLoss returns if the passed in error is an ErrDataLoss
func IsDataLoss(err error) bool {
	_, ok := getImplementer(err).(ErrDataLoss)
	return ok
}
