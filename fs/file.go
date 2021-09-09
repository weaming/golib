package fs

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func Exist(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func IsFile(path string) bool {
	return fileMode(path) == "file"
}

func IsDir(path string) bool {
	return fileMode(path) == "directory"
}

func fileMode(path string) string {
	fi, err := os.Stat(path)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	switch mode := fi.Mode(); {
	case mode.IsDir():
		return "directory"
	case mode.IsRegular():
		return "file"
	default:
		return ""
	}
}

func CreateDirIfNotExist(dir string) error {
	mode := os.FileMode(0777)
	_, err := os.Stat(dir)
	if os.IsNotExist(err) {
		if err := os.MkdirAll(dir, mode); err != nil {
			return err
		}
	}
	return nil
}

func PrepareDirFor(path string) {
	e := CreateDirIfNotExist(filepath.Dir(path))
	if e != nil {
		panic(e)
	}
}

func OpenFile(path string) (*os.File, error) {
	path = ExpandUser(path)
	dir := filepath.Dir(path)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0777)
		if err != nil {
			return nil, err
		}
	}
	return os.Create(path)
}

func OpenFileForAppend(path string) (*os.File, error) {
	return os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
}

func ReadFile(path string) []byte {
	if IsFile(path) {
		dat, err := ioutil.ReadFile(path)
		if err != nil {
			log.Println(err)
			return nil
		}
		return dat
	}
	return nil
}

func WriteFile(path string, content []byte) {
	if e := ioutil.WriteFile(path, content, 0644); e != nil {
		log.Fatal(e)
	}
}

func AppendFile(path string, line string) error {
	PrepareDirFor(path)
	f, e := OpenFileForAppend(path)
	defer f.Close()
	if e != nil {
		return e
	}

	if !strings.HasSuffix(line, "\n") {
		line += "\n"
	}
	_, e = f.WriteString(line)
	return e
}

func ItemsTrimed(str string, sep string) []string {
	xs := []string{}
	for _, x := range strings.Split(str, sep) {
		xs = append(xs, strings.TrimSpace(x))
	}
	return xs
}
