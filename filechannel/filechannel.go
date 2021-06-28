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
	ch   chan string
	File string
}

func NewFileChan(path string, size int) *FileChan {
	fc := &FileChan{}
	fc.ch = make(chan string, size)
	fc.File = path
	return fc
}

func (r *FileChan) In(x string) {
	r.Lock()
	defer r.Unlock()
	r.ch <- x
	e := fs.AppendFile(r.File, x)
	if e != nil {
		log.Println(`write file "%v" err: %v`, r.File, e)
	}
}

func (r *FileChan) Out() string {
	r.Lock()
	defer r.Unlock()
	x := <-r.ch
	exec.Exec(fmt.Sprintf("sed '0,/^%v$/{//d}' %v -i", x, r.File))
	return x
}
