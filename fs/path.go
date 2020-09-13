package fs

import (
	"os"
	"os/user"
	"strings"
)

func ExpandUser(s string) string {
	s = os.ExpandEnv(s)
	if strings.HasPrefix(s, "~/") {
		usr, _ := user.Current()
		home := usr.HomeDir
		s = strings.Replace(s, "~", home, 1)
	}
	return s
}
