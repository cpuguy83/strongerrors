package errclass

import (
	"errors"
	"testing"
)

var errTest = errors.New("this is a test")

func TestNotFound(t *testing.T) {
	e := NotFound(errTest)
	if !IsNotFound(e) {
		t.Fatalf("expected not found error, got: %T", e)
	}
}

func TestConflict(t *testing.T) {
	e := Conflict(errTest)
	if !IsConflict(e) {
		t.Fatalf("expected conflcit error, got: %T", e)
	}
}

func TestForbidden(t *testing.T) {
	e := Forbidden(errTest)
	if !IsForbidden(e) {
		t.Fatalf("expected forbidden error, got: %T", e)
	}
}

func TestInvalidArgument(t *testing.T) {
	e := InvalidArgument(errTest)
	if !IsInvalidArgument(e) {
		t.Fatalf("expected invalid argument error, got %T", e)
	}
}

func TestNotImplemented(t *testing.T) {
	e := NotImplemented(errTest)
	if !IsNotImplemented(e) {
		t.Fatalf("expected not implemented error, got %T", e)
	}
}

func TestNotModified(t *testing.T) {
	e := NotModified(errTest)
	if !IsNotModified(e) {
		t.Fatalf("expected not modified error, got %T", e)
	}
}

func TestUnauthorized(t *testing.T) {
	e := Unauthorized(errTest)
	if !IsUnauthorized(e) {
		t.Fatalf("expected unauthorized error, got %T", e)
	}
}

func TestUnknown(t *testing.T) {
	e := Unknown(errTest)
	if !IsUnknown(e) {
		t.Fatalf("expected unknown error, got %T", e)
	}
}
