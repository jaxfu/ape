package internal

import (
	"fmt"

	"github.com/jaxfu/ape/compiler/internal/assembler"
	"github.com/jaxfu/ape/compiler/internal/parser"
	"github.com/jaxfu/ape/compiler/internal/preprocessor"
	"github.com/jaxfu/ape/compiler/internal/scanner"
	"github.com/jaxfu/ape/components"
)

type Compiler struct{}

func DefaultCompiler() Compiler {
	return Compiler{}
}

func (c Compiler) File(path string, bytes []byte) (components.Components, error) {
	// preprocess
	rawComp, err := preprocessor.NewPreprocessor().File(path, bytes)
	if err != nil {
		return components.Components{}, fmt.Errorf("Preprocessor.File: %+v", err)
	}

	// scan
	scanned, err := scanner.NewScanner().ScanComponent(rawComp)
	if err != nil {
		return components.Components{}, fmt.Errorf("Scanner.ScanComponent: %+v", err)
	}

	// parse
	parsed, err := parser.NewParser().ParseObject(scanned, true)
	if err != nil {
		return components.Components{}, fmt.Errorf("Parser.ParseRoute: %+v", err)
	}

	// assemble
	assembled, err := assembler.NewAssembler().AssembleObject(parsed)
	if err != nil {
		return components.Components{}, fmt.Errorf("Assembler.AssembleRoute: %+v", err)
	}

	ac := components.Components{
		assembled.ComponentId: assembled,
	}

	return ac, nil
}
