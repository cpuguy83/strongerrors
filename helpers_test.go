package errclass

import (
	"errors"
	"testing"
)

var errTest = errors.New("this is a test")

type causal interface {
	Cause() error
}

func TestNotFound(t *testing.T) {
	e := NotFound(errTest)
	if !IsNotFound(e) {
		t.Fatalf("expected not found error, got: %T", e)
	}
	if cause := e.(causal).Cause(); cause != errTest {
		t.Fatalf("causual should be errTest, got: %v", cause)
	}
}

func TestConflict(t *testing.T) {
	e := Conflict(errTest)
	if !IsConflict(e) {
		t.Fatalf("expected conflcit error, got: %T", e)
	}
	if cause := e.(causal).Cause(); cause != errTest {
		t.Fatalf("causual should be errTest, got: %v", cause)
	}
}

func TestForbidden(t *testing.T) {
	e := Forbidden(errTest)
	if !IsForbidden(e) {
		t.Fatalf("expected forbidden error, got: %T", e)
	}
	if cause := e.(causal).Cause(); cause != errTest {
		t.Fatalf("causual should be errTest, got: %v", cause)
	}
}

func TestInvalidArgument(t *testing.T) {
	e := InvalidArgument(errTest)
	if !IsInvalidArgument(e) {
		t.Fatalf("expected invalid argument error, got %T", e)
	}
	if cause := e.(causal).Cause(); cause != errTest {
		t.Fatalf("causual should be errTest, got: %v", cause)
	}
}

func TestNotImplemented(t *testing.T) {
	e := NotImplemented(errTest)
	if !IsNotImplemented(e) {
		t.Fatalf("expected not implemented error, got %T", e)
	}
	if cause := e.(causal).Cause(); cause != errTest {
		t.Fatalf("causual should be errTest, got: %v", cause)
	}
}

func TestNotModified(t *testing.T) {
	e := NotModified(errTest)
	if !IsNotModified(e) {
		t.Fatalf("expected not modified error, got %T", e)
	}
	if cause := e.(causal).Cause(); cause != errTest {
		t.Fatalf("causual should be errTest, got: %v", cause)
	}
}

func TestAlreadyExists(t *testing.T) {
	e := AlreadyExists(errTest)
	if !IsAlreadyExists(e) {
		t.Fatalf("expected already exists error, got %T", e)
	}
	if cause := e.(causal).Cause(); cause != errTest {
		t.Fatalf("causual should be errTest, got: %v", cause)
	}
}

func TestUnauthorized(t *testing.T) {
	e := Unauthorized(errTest)
	if !IsUnauthorized(e) {
		t.Fatalf("expected unauthorized error, got %T", e)
	}
	if cause := e.(causal).Cause(); cause != errTest {
		t.Fatalf("causual should be errTest, got: %v", cause)
	}
}

func TestUnauthenticated(t *testing.T) {
	e := Unauthenticated(errTest)
	if !IsUnauthenticated(e) {
		t.Fatalf("expected unauthenticated error, got %T", e)
	}
	if cause := e.(causal).Cause(); cause != errTest {
		t.Fatalf("causual should be errTest, got: %v", cause)
	}
}

func TestUnknown(t *testing.T) {
	e := Unknown(errTest)
	if !IsUnknown(e) {
		t.Fatalf("expected unknown error, got %T", e)
	}
	if cause := e.(causal).Cause(); cause != errTest {
		t.Fatalf("causual should be errTest, got: %v", cause)
	}
}

func TestCancelled(t *testing.T) {
	e := Cancelled(errTest)
	if !IsCancelled(e) {
		t.Fatalf("expected canclled error, got %T", e)
	}
	if cause := e.(causal).Cause(); cause != errTest {
		t.Fatalf("causual should be errTest, got: %v", cause)
	}
}

func TestDeadline(t *testing.T) {
	e := Deadline(errTest)
	if !IsDeadline(e) {
		t.Fatalf("expected deadline error, got %T", e)
	}
	if cause := e.(causal).Cause(); cause != errTest {
		t.Fatalf("causual should be errTest, got: %v", cause)
	}
}

func TestExhausted(t *testing.T) {
	e := Exhausted(errTest)
	if !IsExhausted(e) {
		t.Fatalf("expected exhausted error, got %T", e)
	}
	if cause := e.(causal).Cause(); cause != errTest {
		t.Fatalf("causual should be errTest, got: %v", cause)
	}
}

func TestIsDataLoss(t *testing.T) {
	e := DataLoss(errTest)
	if !IsDataLoss(e) {
		t.Fatalf("expected data loss error, got %T", e)
	}
	if cause := e.(causal).Cause(); cause != errTest {
		t.Fatalf("causual should be errTest, got: %v", cause)
	}
}
