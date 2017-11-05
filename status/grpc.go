package status

import (
	"github.com/cpuguy83/errclass"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// FromGRPC returns an error class from the provided GPRC status
// If the status is nil or OK, this will return nil
// nolint: gocyclo
func FromGRPC(s *status.Status) error {
	if s == nil || s.Code() == codes.OK {
		return nil
	}

	switch s.Code() {
	case codes.InvalidArgument:
		return errclass.InvalidArgument(s.Err())
	case codes.NotFound:
		return errclass.NotFound(s.Err())
	case codes.Unimplemented:
		return errclass.NotImplemented(s.Err())
	case codes.DeadlineExceeded:
		return errclass.Deadline(s.Err())
	case codes.Canceled:
		return errclass.Cancelled(s.Err())
	case codes.AlreadyExists:
		return errclass.AlreadyExists(s.Err())
	case codes.PermissionDenied:
		return errclass.Unauthorized(s.Err())
	case codes.Unauthenticated:
		return errclass.Unauthenticated(s.Err())
	// TODO(cpuguy83): consider more granular errors for these cases
	case codes.FailedPrecondition, codes.Aborted, codes.Unavailable, codes.OutOfRange:
		return errclass.Conflict(s.Err())
	case codes.ResourceExhausted:
		return errclass.Exhausted(s.Err())
	case codes.DataLoss:
		return errclass.DataLoss(s.Err())
	default:
		return errclass.Unknown(s.Err())
	}
}

// ToGRPC takes the passed in error and converts it to a GRPC status error
// If the passed in error is already a gprc status error, then it is returned unmodified
// If the passed in error is nil, then a nil error is returned.
// nolint: gocyclo
func ToGRPC(err error) error {
	if _, ok := status.FromError(err); ok {
		return err
	}

	switch {
	case errclass.IsNotFound(err):
		return status.Error(codes.NotFound, err.Error())
	case errclass.IsConflict(err), errclass.IsNotModified(err):
		return status.Error(codes.FailedPrecondition, err.Error())
	case errclass.IsInvalidArgument(err):
		return status.Error(codes.InvalidArgument, err.Error())
	case errclass.IsAlreadyExists(err):
		return status.Error(codes.AlreadyExists, err.Error())
	case errclass.IsCancelled(err):
		return status.Error(codes.Canceled, err.Error())
	case errclass.IsDeadline(err):
		return status.Error(codes.DeadlineExceeded, err.Error())
	case errclass.IsUnauthorized(err):
		return status.Error(codes.PermissionDenied, err.Error())
	case errclass.IsUnauthenticated(err):
		return status.Error(codes.Unauthenticated, err.Error())
	case errclass.IsForbidden(err), errclass.IsNotImplemented(err):
		return status.Error(codes.Unimplemented, err.Error())
	case errclass.IsExhausted(err):
		return status.Error(codes.ResourceExhausted, err.Error())
	case errclass.IsDataLoss(err):
		return status.Error(codes.DataLoss, err.Error())
	case errclass.IsSystem(err):
		return status.Error(codes.Internal, err.Error())
	case errclass.IsUnavailable(err):
		return status.Error(codes.Unavailable, err.Error())
	default:
		return status.Error(codes.Unknown, err.Error())
	}
}
