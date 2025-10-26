package filechannel

import (
	"bufio"
	"log"
	"os"
	"strings"
	"sync"

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
		log.Printf(`write file "%v" err: %v`, r.File, e)
	}
}

func (r *FileChan) Has(x string) bool {
	r.Lock()
	defer r.Unlock()

	file, e := os.Open(r.File)
	if e != nil {
		return false
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == x {
			return true
		}
	}

	return false
}

func (r *FileChan) Out() string {
	x := <-r.ch
	r.OutByValue(x)
	return x
}

func (r *FileChan) OutByValue(x string) {
	r.Lock()
	defer r.Unlock()

	file, e := os.Open(r.File)
	if e != nil {
		log.Printf(`read file "%v" err: %v`, r.File, e)
		return
	}

	var lines []string
	scanner := bufio.NewScanner(file)
	foundFirst := false

	for scanner.Scan() {
		line := scanner.Text()
		if line == x {
			if r.allowSame {
				if !foundFirst {
					foundFirst = true
					continue
				}
			} else {
				continue
			}
		}
		lines = append(lines, line)
	}
	file.Close()

	if e := scanner.Err(); e != nil {
		log.Printf(`scan file "%v" err: %v`, r.File, e)
		return
	}

	content := strings.Join(lines, "\n")
	if len(lines) > 0 {
		content += "\n"
	}

	e = os.WriteFile(r.File, []byte(content), 0644)
	if e != nil {
		log.Printf(`write file "%v" err: %v`, r.File, e)
	}
}
