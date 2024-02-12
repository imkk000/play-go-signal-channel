package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	defer fmt.Println("stopped")

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGTERM, syscall.SIGINT)
	defer close(sig)
	<-sig

	fmt.Println("stopping")
	time.Sleep(1 * time.Second)
}
