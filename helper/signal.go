package helper

import (
	"log"
	"os"
	"os/signal"
)

func CaptureInterrupt() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	<-interrupt
	log.Fatal("Interrupt")
}
