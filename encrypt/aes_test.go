package encrypt

import (
	"crypto/sha256"
	"fmt"
	"testing"
)

func Sha256(content []byte) string {
	return fmt.Sprintf("%x", sha256.Sum256(content))
}

func TestAes(t *testing.T) {
	key := []byte(Sha256([]byte("password"))[:32])
	from := "very very long text"
	if AESDecrypt(key, AESEncrypt(key, from)) != from {
		t.Fail()
	}
}
