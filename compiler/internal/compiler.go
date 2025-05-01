package internal

import (
	"fmt"

	"github.com/jaxfu/ape/components"
)

type Compiler struct{}

func DefaultCompiler() Compiler {
	return Compiler{}
}

var componentTypes = components.ComponentTypes.Types()

// just filepath and read file?
// would allow to use compiler as cli
func (c Compiler) File(path string, bytes []byte) (
	components.Components,
	error,
) {
	// lines := getLines(string(bytes))
	// comps := getSourceMapComponents(lines)
	// dev.PrettyPrint(comps)

	_, err := Lex(string(bytes))
	if err != nil {
		return nil, fmt.Errorf(
			"error lexing file '%s': %+v",
			path,
			err,
		)
	}
	// dev.PrettyPrint(tokens)

	return components.Components{}, nil
}
