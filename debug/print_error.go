package debug

import (
	"log"
	"os"
)

func Debug(v interface{}) {
	if os.Getenv("DEBUG") != "" {
		log.Println(v)
	}
}
