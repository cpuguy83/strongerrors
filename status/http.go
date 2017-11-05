package status

import (
	"net/http"

	"github.com/cpuguy83/errclass"
)

// HTTPCode takes an error and returns the HTTP status code for the given error
// If a match is found then the second return argument will be true, otherwise it will be false.
// nolint: gocyclo
func HTTPCode(err error) (int, bool) {
	switch {
	case errclass.IsNotFound(err):
		return http.StatusNotFound, true
	case errclass.IsInvalidArgument(err):
		return http.StatusBadRequest, true
	case errclass.IsConflict(err):
		return http.StatusConflict, true
	case errclass.IsUnauthenticated(err), errclass.IsForbidden(err):
		return http.StatusForbidden, true
	case errclass.IsUnauthorized(err):
		return http.StatusUnauthorized, true
	case errclass.IsUnavailable(err):
		return http.StatusServiceUnavailable, true
	case errclass.IsForbidden(err):
		return http.StatusForbidden, true
	case errclass.IsAlreadyExists(err), errclass.IsNotModified(err):
		return http.StatusNotModified, true
	case errclass.IsNotImplemented(err):
		return http.StatusNotImplemented, true
	case errclass.IsSystem(err) || errclass.IsUnknown(err) || errclass.IsDataLoss(err) || errclass.IsExhausted(err):
		return http.StatusInternalServerError, true
	default:
		return http.StatusInternalServerError, false
	}
}
