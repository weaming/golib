package fs

import (
	"os/user"
	"strings"
)

func ExpandUser(s string) string {
	usr, _ := user.Current()
	home := usr.HomeDir
	if strings.HasPrefix(s, "$HOME") {
		s = strings.Replace(s, "$HOME", home, 1)
	}
	if strings.HasPrefix(s, "~/") {
		s = strings.Replace(s, "~", home, 1)
	}
	return s
}
