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
		if err := server.Serve(); err != nil {
			log.Fatalf("Serve err: %v\n", err)
		}
	}()
	go func() {
		defer wg.Done()
		sig := make(chan os.Signal, 1)
		signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
		<-sig
		if err := server.Shutdown(); err != nil {
			log.Printf("Shutdown err: %v\n", err)
		}
	}()
	wg.Wait()
	log.Println("bye")
}
