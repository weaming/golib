package exec

import (
	"io"
	"log"
	"os/exec"

	"github.com/weaming/golib/fs"
)

func ExecGetOutput(command string, cmdChan chan<- *exec.Cmd, logfile string) (string, error) {
	cmd := exec.Command("bash", []string{"-c", command}...)

	var r_tee io.Reader
	if logfile != "" {
		// PrepareDirFor(logfile)
		log, e := fs.OpenFileForAppend(logfile)
		if e != nil {
			return "", e
		}

		r, w := io.Pipe()
		r_tee = io.TeeReader(r, log)
		cmd.Stdout = w
		cmd.Stderr = w
	}

	// log.Println("cmd:", cmd)
	if cmdChan != nil {
		cmdChan <- cmd
	}

	if logfile != "" {
		errChan := make(chan error, 1)
		// run in another goroutine
		go func() {
			e := cmd.Run()
			errChan <- e
		}()

		r := r_tee

		// copy from ioutil.ReadAll
		ReadAll := func() ([]byte, error) {
			step := 512
			// step := 64
			b := make([]byte, 0, step)
			for {
				select {
				case resultErr := <-errChan:
					return b, resultErr
				default:
					if len(b) == cap(b) {
						b = append(b, 0)[:len(b)]
					}
					n, err := r.Read(b[len(b):minInt(len(b)+step, cap(b))])
					b = b[:len(b)+n]
					if err != nil {
						if err == io.EOF {
							err = nil
						}
						return b, err
					}
				}
			}
		}

		bs, e := ReadAll()
		return string(bs), e
	} else {
		stdoutStderr, err := cmd.CombinedOutput()
		return string(stdoutStderr), err
	}
}

func Exec(command string) {
	cmd := exec.Command("bash", []string{"-c", command}...)
	stdoutStderr, err := cmd.CombinedOutput()
	HandlerResult(command, string(stdoutStderr), err)
}

func HandlerResult(command string, output string, err error) {
	if err != nil {
		log.Printf("execute command `%v` error: %v", command, err)
	}
	log.Printf("execute command `%v` output: %v", command, string(output))
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}
