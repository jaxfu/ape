package preprocessor

import (
	"github.com/jaxfu/ape/compiler/internal/preprocessor/internal"
)

type Preprocessor interface {
	File(string, []byte) (RawComponent, error)
}

func NewPreprocessor() Preprocessor {
	return internal.DefaultPreprocessor()
}

type (
	RawComponent = internal.RawComponent
)
