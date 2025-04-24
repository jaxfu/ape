package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	corepkg "github.com/jaxfu/ape/engine/core"
	"github.com/jaxfu/ape/engine/pkg/dev"
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
	defer core.Db.Close()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	shutdownErr := make(chan error, 1)
	serverClosedChan := make(chan bool)
	go func() {
		shutdownErr <- core.Server.Start(ctx)
		serverClosedChan <- true
	}()

	select {
	case sig := <-sigChan:
		fmt.Printf("\nReceived shutdown signal: %s, shutting down...\n", sig)
		cancel()
	case err := <-shutdownErr:
		fmt.Printf("\nServer error, stopping: %v\n", err)
	}

	<-serverClosedChan
	dev.Shutdown()
}
