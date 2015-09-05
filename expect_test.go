package algorithms

import (
	"fmt"
	"path"
	"reflect"
	"runtime"
	"testing"
)

func Expect(t *testing.T) *expectation {
	return &expectation{t: t}
}

type expectation struct {
	t *testing.T
}

func (e *expectation) Equal(a, b interface{}) {
	if !reflect.DeepEqual(a, b) {
		e.t.Error(e.originalCaller(), a, "does not equal", b)
	}
}

func (e *expectation) originalCaller() string {
	_, file, number, ok := runtime.Caller(2)
	if !ok {
		panic("can't get caller")
	}
	return fmt.Sprintf("%s:%d:", path.Base(file), number)
}
