package http

import (
	"testing"
)

func TestUrl_encoder(t *testing.T) {
	if _, err := URLEncoded("test"); err != nil {
		t.Fail()
	}

	if _, err := URLEncoded("$#%B^&M*(>(*)>(*?(?```````HGKKJ:"); err == nil {
		t.Fail()
	}
}
