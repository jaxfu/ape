package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	corepkg "github.com/jaxfu/ape/engine/core"
	"github.com/jaxfu/ape/engine/core/bus"
	"github.com/jaxfu/ape/engine/pkg/dev"
	"github.com/jaxfu/ape/engine/pkg/extras"
)

const (
	TEST_ROOT_DIR string = "../example"
)

func handleEvents(bus *bus.Bus) {
	for event := range bus.Events {
		fmt.Println("event: ")
		fmt.Println(event)
	}
}

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

	go handleEvents(core.Bus)

	select {
	case sig := <-sigChan:
		fmt.Printf("\nReceived shutdown signal: %s, shutting down...\n", sig)
		cancel()
	case err := <-shutdownErr:
		fmt.Printf("\nServer error, stopping: %v\n", err)
	}

	<-serverClosedChan
	dev.Shutdown()

	// rootDir, err := filepath.Abs(TEST_ROOT_DIR)
	// if err != nil {
	// 	log.Printf("illegal filepath '%s': %+v\n", TEST_ROOT_DIR, err)
	// }
	//
	// // read file
	// filehandler := filehandler.NewFileHandler()
	// rawFile, err := filehandler.ReadFile(
	// 	fmt.Sprintf("%s/objects/Todo.toml", rootDir),
	// )
	// if err != nil {
	// 	log.Fatalf("FileHandler.ReadFile: %+v", err)
	// }
	//
	// compiled, err := compiler.NewCompiler().File(
	// 	rawFile.Path(),
	// 	rawFile.Bytes(),
	// )
	// if err != nil {
	// 	log.Fatalf("error compiling file %s: %+v", rawFile.Path(), err)
	// }
	// // dev.PrettyPrint(compiled)
	//
	// str := store.NewStore()
	// lnkr := linker.NewLinker(str)
	// ac, err := lnkr.LinkAll(compiled)
	// if err != nil {
	// 	log.Fatalf("Linker.LinkAll: %+v", err)
	// }
	// dev.PrettyPrint(ac)
}
