package fs

import (
	"fmt"
	"os"
	"path/filepath"
)

func Exist(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	} else {
		return true
	}
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
		// do directory stuff
		return "directory"
	case mode.IsRegular():
		// do file stuff
		return "file"
	default:
		return ""
	}
}

func CreateDirIfNotExist(dir string, force bool) error {
	mode := os.FileMode(0777)
	fi, err := os.Stat(dir)
	if os.IsNotExist(err) {
		// prepare dir
		if err := os.MkdirAll(dir, mode); err != nil {
			return err
		}
	} else {
		// if is normal file
		if force && fi.Mode().IsRegular() {
			if err := os.Remove(dir); err != nil {
				return err
			}
			if err := os.MkdirAll(dir, mode); err != nil {
				return err
			}
		}
	}
	return nil
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
