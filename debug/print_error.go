package debug

import (
	"log"
	"os"
)

func Debug(e interface{}) {
	if os.Getenv("DEBUG") != "" {
		log.Println(e)
	}
}
