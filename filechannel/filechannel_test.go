package filechannel

import (
	"fmt"
	"os"
	"testing"

	"github.com/weaming/golib/fs"
)

func TestFileChannel(t *testing.T) {
	testFile := "./set"
	os.Remove(testFile)
	defer os.Remove(testFile)

	fc := NewFileChan(testFile, 100, true)
	fs.AppendFile(testFile, "a")
	fs.AppendFile(testFile, "b")

	fc.In("a")
	fc.In("b")
	fc.In("c")

	fc.In("a")
	fc.In("b")
	fc.In("c")

	fmt.Println(fc.Out())
	fmt.Println(fc.Out())
	fmt.Println(fc.Out())
	fmt.Println(fc.Out())
	fmt.Println("-----------")
	if string(fs.ReadFile(testFile)) != `b
a
b
c
` {
		t.Fatal("fail")
	}
}
