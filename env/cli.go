package env

import (
	"fmt"
	"os"
)

func IsHelp() bool {
	last := os.Args[len(os.Args)-1]
	return last == "-h" || last == "--help"
}

func PrintHelp(help string) {
	if IsHelp() {
		fmt.Println(help)
		os.Exit(0)
	}
}
