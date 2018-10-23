package gorandstring

import (
	"testing"
)

func TestGet(t *testing.T) {
	v := Get(63)
	if len(v) != 63 {
		t.Errorf("unexpected len, exp [%d] got [%d]", 63, len(v))
	}
	v2 := Get(63)
	if v == v2 {
		t.Error("same values")
	}
}
