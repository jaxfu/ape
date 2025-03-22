package scanner

import (
	"github.com/jaxfu/ape/engine/compiler/internal/preprocessor"
	"github.com/jaxfu/ape/engine/compiler/internal/scanner/internal"
)

type Scanner interface {
	ScanComponent(preprocessor.RawComponent) (ScannedComponent, error)
}

func NewScanner() Scanner {
	return internal.DefaultScanner()
}

type (
	ScannedComponent = internal.ScannedComponent
)
