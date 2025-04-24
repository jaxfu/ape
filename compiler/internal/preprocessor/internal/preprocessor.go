package internal

import (
	"fmt"
	"strings"

	"github.com/jaxfu/ape/components"
)

type Preprocessor struct{}

func DefaultPreprocessor() Preprocessor {
	return Preprocessor{}
}

type RawFile interface {
	Path() string
	Bytes() []byte
}

type RawFileMetadata struct {
	Path string `json:"filepath" toml:"filepath"`
}
type RawComponent struct {
	ComponentType components.ComponentType `json:"component_type" toml:"component_type"`
	IsFromFile   bool            `json:"is_file" toml:"is_file"`
	FileMetadata RawFileMetadata `json:"metadata" toml:"metadata"`
	Bytes        []byte          `json:"bytes" toml:"bytes"`
}

func (pr Preprocessor) File(path string, bytes []byte) (RawComponent, error) {
	rawComp := RawComponent{
		ComponentType: "",
		IsFromFile:    true,
		Bytes:         bytes,
		FileMetadata: RawFileMetadata{
			Path: path,
		},
	}

	// TODO: better way to get types from files
	compTypesMap := map[string]components.ComponentType{
		"props":   components.COMPONENT_TYPE_PROP,
		"objects": components.COMPONENT_TYPE_OBJECT,
		"routes":  components.COMPONENT_TYPE_ROUTE,
	}

	cType := ""
	for k, v := range compTypesMap {
		if strings.Contains(strings.ToLower(path), k) {
			cType = v
		}
	}
	if cType == "" {
		return RawComponent{}, fmt.Errorf(
			"could not determine type for file %s",
			path,
		)
	}
	rawComp.ComponentType = cType

	return rawComp, nil
}
