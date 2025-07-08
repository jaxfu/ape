package internal

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/jaxfu/ape/compiler/internal/assembler"
	lexerPkg "github.com/jaxfu/ape/compiler/internal/lexer"
	parserPkg "github.com/jaxfu/ape/compiler/internal/parser"
	"github.com/jaxfu/ape/compiler/internal/shared"
	"github.com/jaxfu/ape/components"
)

const (
	TEST_DIR string = "../test"
)

var (
	tokens []shared.Token
	ast    shared.Ast
	comps  components.ComponentMap
	err    error
)

func TestCompiler(t *testing.T) {
	t.Run("Lexer", func(t *testing.T) {
		testLexer(t)
	})
	// marsh, _ := json.MarshalIndent(tokens, "", " ")
	// t.Log(string(marsh))

	t.Run("Parser", func(t *testing.T) {
		testParser(t)
	})
	// marsh, _ := json.MarshalIndent(ast, "", " ")
	// t.Log(string(marsh))

	t.Run("Assembler", func(t *testing.T) {
		testAssembler(t)
	})
	marsh, _ := json.MarshalIndent(comps, "", " ")
	t.Log(string(marsh))
}

func testLexer(t *testing.T) {
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
		t.Errorf("lexer.Lex on file '%s':\n%+v", path, err)
	}
}

func testParser(t *testing.T) {
	ast, _, err = parserPkg.Parse(tokens, PREALLOC)
	if err != nil {
		t.Fatalf("parser.Parse:\n%+v\n", err)
	}
}

func testAssembler(t *testing.T) {
	comps, err = assembler.Assemble(ast)
	if err != nil {
		t.Errorf("assembler.Assemble:\n%+v\n", err)
	}
}
