package status

import (
	"net/http"

	"github.com/cpuguy83/strongerrors"
)

// HTTPCode takes an error and returns the HTTP status code for the given error
// If a match is found then the second return argument will be true, otherwise it will be false.
// nolint: gocyclo
func HTTPCode(err error) (int, bool) {
	switch {
	case strongerrors.IsNotFound(err):
		return http.StatusNotFound, true
	case strongerrors.IsInvalidArgument(err):
		return http.StatusBadRequest, true
	case strongerrors.IsConflict(err):
		return http.StatusConflict, true
	case strongerrors.IsUnauthenticated(err), strongerrors.IsForbidden(err):
		return http.StatusForbidden, true
	case strongerrors.IsUnauthorized(err):
		return http.StatusUnauthorized, true
	case strongerrors.IsUnavailable(err):
		return http.StatusServiceUnavailable, true
	case strongerrors.IsForbidden(err):
		return http.StatusForbidden, true
	case strongerrors.IsAlreadyExists(err), strongerrors.IsNotModified(err):
		return http.StatusNotModified, true
	case strongerrors.IsNotImplemented(err):
		return http.StatusNotImplemented, true
	case strongerrors.IsSystem(err) || strongerrors.IsUnknown(err) || strongerrors.IsDataLoss(err) || strongerrors.IsExhausted(err):
		return http.StatusInternalServerError, true
	default:
		return http.StatusInternalServerError, false
	}
}

// FromHTTPCode wraps the passed in error in a strong error based on the passed in
// http status code.
//
// If the http status code is < 300, the original error is returned.
// If the passed in error is nil, the returned error will be nil.
// If the status code does not match a known code the error will be wrapped in
// an `Unknown` error.
func FromHTTPResponse(httpCode int, err error) error {
	if httpCode < 300 || err == nil {
		return err
	}

	switch httpCode {
	case http.StatusNotFound:
		return strongerrors.NotFound(err)
	case http.StatusBadRequest:
		return strongerrors.InvalidParameter(err)
	case http.StatusConflict:
		return strongerrors.Conflict(err)
	case http.StatusForbidden:
		return strongerrors.Forbidden(err)
	case http.StatusUnauthorized:
		return strongerrors.Unauthorized(err)
	case http.StatusUnavailable:
		return strongerrors.Unavailable(err)
	case http.StatusNotModified:
		return strongerrors.NotModified(err)
	case http.StatusNotImplemented:
		return strongerrors.NotImplemented(err)
	default:
		return strongerrors.Unknown(err)
	}
}
