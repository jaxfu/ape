package main

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/jaxfu/ape/engine/compiler"
	"github.com/jaxfu/ape/engine/core/linker"
	"github.com/jaxfu/ape/engine/core/store"
	"github.com/jaxfu/ape/engine/pkg/dev"
	"github.com/jaxfu/ape/engine/pkg/extras"
	"github.com/jaxfu/ape/engine/pkg/filehandler"
	_ "modernc.org/sqlite"
)

const (
	DB_NAME       string = "ape.db"
	INIT_DB_SQL   string = "core/db/sql/schemas.sql"
	TEST_ROOT_DIR string = "../example"
	BASE_URL      string = "localhost"
	PORT          uint   = 5000
	CLIENT_DIR    string = "clients/web/dist"
)

func main() {
	fmt.Println(extras.SPLASH)

	rootDir, err := filepath.Abs(TEST_ROOT_DIR)
	if err != nil {
		log.Printf("illegal filepath '%s': %+v\n", TEST_ROOT_DIR, err)
	}

	// read file
	filehandler := filehandler.NewFileHandler()
	rawFile, err := filehandler.ReadFile(
		fmt.Sprintf("%s/objects/Todo.toml", rootDir),
	)
	if err != nil {
		log.Fatalf("FileHandler.ReadFile: %+v", err)
	}

	compiled, err := compiler.NewCompiler().File(
		rawFile.Path(),
		rawFile.Bytes(),
	)
	if err != nil {
		log.Fatalf("error compiling file %s: %+v", rawFile.Path(), err)
	}
	// dev.PrettyPrint(compiled)

	str := store.NewStore()
	lnkr := linker.NewLinker(str)
	ac, err := lnkr.LinkAll(compiled)
	if err != nil {
		log.Fatalf("Linker.LinkAll: %+v", err)
	}
	dev.PrettyPrint(ac)

	// rte, err := store.Get[apeComponents.Route](
	// 	str.Components,
	// 	"Todos.routes.Create",
	// )
	// if err != nil {
	// 	log.Fatalf("Store.Get: %+v", err)
	// }
	// dev.PrettyPrint(rte)

	// db, err := db.NewDb(DB_NAME, INIT_DB_SQL)
	// if err != nil {
	// 	log.Fatalf("error opening db at %s: %+v\n", DB_NAME, err)
	// }
	// defer db.Conn().Close()
	//
	// m, err := json.Marshal(compiled.Routes[0])
	// if err != nil {
	// 	log.Fatalf("json.Marshal: %+v", err)
	// }
	// if err := db.InsertComponent("test", m); err != nil {
	// 	log.Fatalf("db.InsertComponent: %+v", err)
	// }
	//
	// ok, data, err := db.GetComponent("test")
	// if !ok {
	// 	log.Fatalf("no component with id %s found", "test")
	// }
	// if err != nil {
	// 	log.Fatalf("db.GetComponent: %+v", err)
	// }
	//
	// rt := components.Route{}
	// if err := json.Unmarshal([]byte(data), &rt); err != nil {
	// 	log.Fatalf("json.Unmarshal: %+v", err)
	// }
	// dev.PrettyPrint(rt.ComponentMetadata)

	// clientDir, err := filepath.Abs(CLIENT_DIR)
	// if err != nil {
	// 	log.Printf("illegal filepath '%s': %+v\n", CLIENT_DIR, err)
	// }
	//
	// core, err := corepkg.InitCore(
	// 	db.Conn(),
	// 	BASE_URL,
	// 	PORT,
	// 	clientDir,
	// )
	// if err != nil {
	// 	log.Fatalf("error starting server: %+v", err)
	// }
	//
	// if err := core.Store.Sync(allComps); err != nil {
	// 	fmt.Printf("error syncing Store: %+v\n\n", err)
	// }
	//
	// ctx, cancel := context.WithCancel(context.Background())
	// defer cancel()
	//
	// sigChan := make(chan os.Signal, 1)
	// signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	//
	// shutdownErr := make(chan error, 1)
	// serverClosedChan := make(chan bool)
	// go func() {
	// 	shutdownErr <- core.Server.StartServer(ctx)
	// 	serverClosedChan <- true
	// }()
	//
	// select {
	// case sig := <-sigChan:
	// 	fmt.Printf("\nReceived shutdown signal: %s, shutting down...\n", sig)
	// 	cancel()
	// case err := <-shutdownErr:
	// 	fmt.Printf("\nServer error, stopping: %v\n", err)
	// }
	//
	// <-serverClosedChan
	// dev.Shutdown()
}
