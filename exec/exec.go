package exec

import (
	"io"
	"log"
	"os/exec"

	"github.com/weaming/golib/fs"
)

func ExecGetOutput(command string, cmdChan chan<- *exec.Cmd, logfile string) (string, error) {
	cmd := exec.Command("bash", []string{"-c", command}...)

	var rTee io.Reader
	var r *io.PipeReader
	var w *io.PipeWriter
	if logfile != "" {
		// PrepareDirFor(logfile)
		log, e := fs.OpenFileForAppend(logfile)
		if e != nil {
			return "", e
		}

		r, w = io.Pipe()
		rTee = io.TeeReader(r, log)
		cmd.Stdout = w
		cmd.Stderr = w
	}

	// log.Println("cmd:", cmd)
	if cmdChan != nil {
		cmdChan <- cmd
	}

	if logfile != "" {
		// copy from ioutil.ReadAll
		ReadAll := func() ([]byte, error) {
			// step := 512
			step := 64 // small step get more instant response in `tail -f`
			b := make([]byte, 0, step)
			doneChan := make(chan bool, 1)
			go func() {
				defer func() { doneChan <- true }()
				for {
					if len(b) == cap(b) {
						b = append(b, 0)[:len(b)]
					}

					n, err := rTee.Read(b[len(b):minInt(len(b)+step, cap(b))])
					b = b[:len(b)+n]
					if err != nil {
						if err == io.EOF {
							err = nil
						}
						return
					}
				}
			}()

			e := cmd.Run()
			w.Close()
			<-doneChan
			return b, e
		}

		bs, e := ReadAll()
		return string(bs), e
	} else {
		stdoutStderr, err := cmd.CombinedOutput()
		return string(stdoutStderr), err
	}
}

func Exec(command string) (string, error) {
	cmd := exec.Command("bash", []string{"-c", command}...)
	stdoutStderr, err := cmd.CombinedOutput()
	out := string(stdoutStderr)
	return out, err
}

func HandlerResult(command string, err error) {
	if err != nil {
		log.Printf("WARN: command `%v` error: %v", command, err)
	} else {
		log.Printf("OK: command `%v` succeed", command)
	}
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}
