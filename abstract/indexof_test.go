package abstract

import (
	"testing"
)

func TestIndexOf(t *testing.T) {
	if IndexOf([]interface{}{"a", "b", "weaming", "d"}, "weaming") < 0 {
		t.Fail()
	}
}
