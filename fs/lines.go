package fs

import (
	"strings"
)

func Lines(text string, fn func(string) bool) (rv []string) {
	text = strings.ReplaceAll(text, "\r\n", "\n")
	xs := strings.Split(strings.TrimSuffix(text, "\n"), "\n")
	return FilterStr(xs, fn)
}

func FilterStr(xs []string, fn func(string) bool) (rv []string) {
	for _, line := range xs {
		if fn(line) {
			rv = append(rv, line)
		}
	}
	return
}

func MapStr(fn func(string) string, xs []string) (rv []string) {
	for _, x := range xs {
		rv = append(rv, fn(x))
	}
	return
}

func LinesWithoutComments(text string) []string {
	return Lines(text, func(s string) bool { return !strings.HasPrefix(strings.TrimSpace(s), "#") })
}

func FileAsLinesWithoutComments(path string) []string {
	bin := ReadFile(path)
	if bin == nil {
		return []string{}
	}
	return LinesWithoutComments(string(bin))
}

func FilterUrlsInFile(path string) []string {
	lines := MapStr(strings.TrimSpace, FileAsLinesWithoutComments(path))
	return FilterStr(lines, func(s string) bool {
		x := strings.HasPrefix(strings.TrimSpace(s), "http")
		return x
	})
}
