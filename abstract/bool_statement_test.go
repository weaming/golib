package abstract

import (
	"testing"
)

func TestBoolStatement(t *testing.T) {
	if OrString("", "x") != "x" {
		t.Fail()
	}
	if AndString("", "x") != "" {
		t.Fail()
	}
	if Or("", "x") != "x" {
		t.Fail()
	}
	if And("", "x") != "" {
		t.Fail()
	}
}
