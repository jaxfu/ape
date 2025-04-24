package internal

import (
	"fmt"
	"github.com/jaxfu/ape/compiler/internal/preprocessor"

	"github.com/BurntSushi/toml"
	"github.com/jaxfu/ape/components"
)

type Scanner struct{}

func DefaultScanner() Scanner {
	return Scanner{}
}

type ScannedComponent struct {
	ComponentType components.ComponentType `json:"component_type" toml:"component_type"`
	Fields        map[string]any           `json:"fields" toml:"fields"`
}

func (s Scanner) ScanComponent(rawComp preprocessor.RawComponent) (ScannedComponent, error) {
	scannedComp := ScannedComponent{
		ComponentType: rawComp.ComponentType,
		Fields:        map[string]any{},
	}

	if err := toml.Unmarshal(rawComp.Bytes, &scannedComp.Fields); err != nil {
		return ScannedComponent{}, fmt.Errorf("toml.Unmarshal: %+v", err)
	}

	return scannedComp, nil
}
