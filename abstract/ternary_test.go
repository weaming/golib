package abstract

import (
	"testing"
)

func TestTernary(t *testing.T) {
	if IfElse("", "a", 3) != 3 {
		t.Fail()
	}
	if IfElse(1, 'a', true) != 'a' {
		t.Fail()
	}
}
