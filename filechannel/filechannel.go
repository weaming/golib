package filechannel

import (
	"fmt"
	"log"
	"sync"

	"github.com/weaming/golib/exec"
	"github.com/weaming/golib/fs"
)

type FileChan struct {
	sync.Mutex
	ch        chan string
	allowSame bool
	File      string
}

// 0 size will not create channel
func NewFileChan(path string, size int, allowSame bool) *FileChan {
	fc := &FileChan{allowSame: allowSame}
	if size > 0 {
		fc.ch = make(chan string, size)
	} else {
		fc.ch = nil
	}
	fc.File = path
	return fc
}

func (r *FileChan) In(x string) {
	r.Lock()
	defer r.Unlock()
	if r.ch != nil {
		r.ch <- x
	}
	e := fs.AppendFile(r.File, x)
	if e != nil {
		log.Println(`write file "%v" err: %v`, r.File, e)
	}
}

func (r *FileChan) Has(x string) bool {
	r.Lock()
	defer r.Unlock()
	_, e := exec.ExecGetOutput(fmt.Sprintf("cat %v | grep -E '^%v$'", r.File, x), nil, "")
	return e == nil
}

func (r *FileChan) Out() string {
	x := <-r.ch
	r.OutByValue(x)
	return x
}

func (r *FileChan) OutByValue(x string) {
	r.Lock()
	defer r.Unlock()
	if r.allowSame {
		// 移除第一个匹配的行
		exec.Exec(fmt.Sprintf("sed '0,/^%v$/{//d}' %v -i", x, r.File))
	} else {
		// 移除匹配的所有行
		exec.Exec(fmt.Sprintf("sed '/^%v$/d' %v -i", x, r.File))
	}
}
