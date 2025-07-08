package internal

import (
	"bufio"
	"fmt"
	"os"

	"github.com/jaxfu/ape/compiler/internal/lexer"
	"github.com/jaxfu/ape/components"
)

const PREALLOC = 1024

type Compiler struct{}

func DefaultCompiler() Compiler {
	return Compiler{}
}

// just filepath and read file?
// would allow to use compiler as cli
func (c Compiler) File(path string, bytes []byte) (
	[]components.Component,
	error,
) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("error opening file '%s': %+v", path, err)
	}
	defer file.Close()
	buf := bufio.NewReader(file)

	lexer := lexer.NewLexer()
	_, err = lexer.Lex(buf, PREALLOC)
	if err != nil {
		return nil, fmt.Errorf(
			"error lexing file '%s': %+v",
			path,
			err,
		)
	}

	return []components.Component{}, nil
}
