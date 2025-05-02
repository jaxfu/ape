package internal

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	lexerPkg "github.com/jaxfu/ape/compiler/internal/lexer"
	parserPkg "github.com/jaxfu/ape/compiler/internal/parser"
	"github.com/jaxfu/ape/compiler/internal/shared"
)

const (
	TEST_DIR string = "../test"
)

var (
	tokens []shared.Token
	ast    shared.Ast
)

func TestCompiler(t *testing.T) {
	t.Run("Lexer", func(t *testing.T) {
		TestLexer(t)
	})

	t.Run("Parser", func(t *testing.T) {
		testParser(t)
	})
}

func TestLexer(t *testing.T) {
	path, err := filepath.Abs(fmt.Sprintf(
		"%s/test.ape",
		TEST_DIR,
	))
	if err != nil {
		t.Errorf("error casting to absolute path: %+v", err)
		return
	}

	file, err := os.Open(path)
	if err != nil {
		t.Errorf("error opening file '%s': %+v", path, err)
	}
	defer file.Close()
	buf := bufio.NewReader(file)

	lexer := lexerPkg.NewLexer()
	tokens, err = lexer.Lex(buf, PREALLOC)
	if err != nil {
		t.Errorf("error lexing file '%s': %+v", path, err)
	}
	// marsh, _ := json.MarshalIndent(tokens, "", " ")
	// t.Log(string(marsh))
}

func testParser(t *testing.T) {
	var err error
	var errs []error
	ast, errs, err = parserPkg.Parse(tokens, PREALLOC)
	if err != nil {
		t.Fatalf("%+v\n", err)
	}

	marsh, _ := json.MarshalIndent(ast, "", " ")
	if len(ast) > 0 {
		t.Log(string(marsh))
	}

	if len(errs) > 0 {
		t.Log("errors: ")
		t.Logf("%+v\n", errs)
	} else {
		t.Log("no errors")
	}
}
