package internal

import (
	"fmt"

	"github.com/jaxfu/ape/engine/compiler/internal/assembler"
	"github.com/jaxfu/ape/engine/compiler/internal/parser"
	"github.com/jaxfu/ape/engine/compiler/internal/preprocessor"
	"github.com/jaxfu/ape/engine/compiler/internal/scanner"
	"github.com/jaxfu/ape/engine/compiler/internal/shared"
)

type Compiler struct{}

func DefaultCompiler() Compiler {
	return Compiler{}
}

func (c Compiler) File(path string, bytes []byte) (shared.CompiledComponents, error) {
	// preprocess
	rawComp, err := preprocessor.NewPreprocessor().File(path, bytes)
	if err != nil {
		return shared.CompiledComponents{}, fmt.Errorf("Preprocessor.File: %+v", err)
	}

	// scan
	scanned, err := scanner.NewScanner().ScanComponent(rawComp)
	if err != nil {
		return shared.CompiledComponents{}, fmt.Errorf("Scanner.ScanComponent: %+v", err)
	}

	// parse
	parsed, err := parser.NewParser().ParseObject(scanned, true)
	if err != nil {
		return shared.CompiledComponents{}, fmt.Errorf("Parser.ParseRoute: %+v", err)
	}

	// assemble
	assembled, err := assembler.NewAssembler().AssembleObject(parsed)
	if err != nil {
		return shared.CompiledComponents{}, fmt.Errorf("Assembler.AssembleRoute: %+v", err)
	}

	return shared.CompiledComponents{
		Props: map[string]shared.CompiledProp{},
		Objects: map[string]shared.CompiledObject{
			assembled.ComponentMetadata.Name: assembled,
		},
		Routes:    map[string]shared.CompiledRoute{},
		Bodies:    map[string]shared.CompiledBody{},
		Requests:  map[string]shared.CompiledRequest{},
		Responses: map[string]shared.CompiledResponse{},
	}, nil
}
