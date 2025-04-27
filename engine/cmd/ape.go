package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	corepkg "github.com/jaxfu/ape/engine/core"
	"github.com/jaxfu/ape/engine/pkg/extras"
)

const (
	TEST_ROOT_DIR string = "../example"
)

func main() {
	fmt.Println(extras.SPLASH)

	core, err := corepkg.InitCore()
	if err != nil {
		log.Fatalf("error starting server: %+v", err)
	}

	coreClosedChan := make(chan bool)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go func() {
		core.Start(ctx)
		coreClosedChan <- true
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	sig := <-sigChan
	fmt.Printf("\nReceived shutdown signal: %s, shutting down...\n", sig)
	cancel()

	<-coreClosedChan
	fmt.Println("Bye!")
}
