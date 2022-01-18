package filechannel

import (
	"fmt"
	"testing"

	"github.com/weaming/golib/fs"
)

func TestFileChannel(t *testing.T) {
	fc := NewFileChan("./set", 100, true)
	fs.AppendFile("./set", "a")
	fs.AppendFile("./set", "b")

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
	if string(fs.ReadFile("./set")) != `b
a
b
c
` {
		t.Fatal("fail")
	}
}
