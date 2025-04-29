package compiler

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

const (
	TEST_DIR = "../example"
)

var compiler = NewCompiler()

func TestCompiler(t *testing.T) {
	path, err := filepath.Abs(fmt.Sprintf("%s/props/Username.toml", TEST_DIR))
	if err != nil {
		t.Errorf("error casting to absolute path: %+v", err)
		return
	}

	bytes, err := os.ReadFile(path)
	if err != nil {
		t.Errorf("error reading file %s: %+v", path, err)
		return
	}

	comps, err := compiler.File(path, bytes)
	if err != nil {
		t.Errorf("error compiling file %s: %+v", path, err)
		return
	}

	marshalled, err := json.MarshalIndent(comps, "", "  ")
	if err != nil {
		t.Errorf("json.MarshalIndent: %+v", err)
		return
	}
	t.Logf("%+v\n", string(marshalled))
}
