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
	TEST_READ_FILE  string = "../test/test.ape"
	TEST_WRITE_FILE string = "testdump.json"
)

var (
	tokens   []shared.Token
	rawNodes []parserPkg.RawNode
	rn       shared.Ast
	comps    components.ComponentsMap
	err      error
)
var marsh []byte = []byte{}

func TestCompiler(t *testing.T) {
	file, err := openFile(TEST_WRITE_FILE)
	if err != nil {
		t.Errorf("error opening file '%s'\n", TEST_WRITE_FILE)
	}
	defer file.Close()

	t.Run("Lexer", func(t *testing.T) {
		testLexer(t)
	})
	// marsh, _ = json.MarshalIndent(tokens, "", " ")
	// fmt.Fprint(file, string(marsh))
	// t.Log(string(marsh))

	t.Run("Parser", func(t *testing.T) {
		testParser(t)
	})
	marsh, _ = json.MarshalIndent(rawNodes, "", " ")
	fmt.Fprint(file, string(marsh))
	// t.Log(string(marsh))

	// t.Run("Assembler", func(t *testing.T) {
	// 	testAssembler(t)
	// })
	// marsh, _ = json.MarshalIndent(comps, "", " ")
	// t.Log(string(marsh))
}

func testLexer(t *testing.T) {
	path, err := filepath.Abs(TEST_READ_FILE)
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
	rawNodes, _, err = parserPkg.Parse(tokens, PREALLOC)
	if err != nil {
		t.Fatalf("parser.Parse:\n%+v\n", err)
	}
}

func testAssembler(t *testing.T) {
	comps, err = assembler.Assemble(rn)
	if err != nil {
		t.Errorf("assembler.Assemble:\n%+v\n", err)
	}
}

func openFile(path string) (*os.File, error) {
	abspath, err := filepath.Abs(path)
	if err != nil {
		return nil, fmt.Errorf("error casting '%s' to absolute path: %+v", path, err)
	}
	file, err := os.OpenFile(
		path,
		os.O_CREATE|os.O_WRONLY,
		0644,
	)
	if err != nil {
		return nil, fmt.Errorf("error opening file '%s'\n", abspath)
	}

	return file, nil
}
