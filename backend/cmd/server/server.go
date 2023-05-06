package main

import (
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/oseau/blockchain-explorer/internal/web"
)

func main() {
	wg := new(sync.WaitGroup)
	wg.Add(2)

	server := web.NewServer()
	go func() {
		defer wg.Done()
		server.Serve()
	}()
	go func() {
		defer wg.Done()
		sig := make(chan os.Signal, 1)
		signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
		<-sig
		server.Shutdown()
	}()
	wg.Wait()
	log.Println("bye")
}
