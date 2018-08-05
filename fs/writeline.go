package fs

import (
	"log"
	"os"
)

func AppendToFile(filepath, line string) {
	f, err := os.OpenFile(filepath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	fatalErr(err)
	defer f.Close()

	// add \n
	if len(line) > 0 {
		if line[len(line)-1] != '\n' {
			line += "\n"
		}
	} else {
		line = "\n"
	}

	if _, err = f.WriteString(line); err != nil {
		panic(err)
	}
}

func fatalErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
