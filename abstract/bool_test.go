package abstract

import (
	"testing"
)

func TestBool(t *testing.T) {
	if !IsTrue(1) {
		t.Fail()
	}
	if IsTrue(0) {
		t.Fail()
	}
	if !IsTrue("x") {
		t.Fail()
	}
	if IsTrue("") {
		t.Fail()
	}
}
