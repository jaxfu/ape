package internal

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestParser(t *testing.T) {
	path, err := filepath.Abs(fmt.Sprintf(
		"%s/test.ape",
		TEST_DIR,
	))
	if err != nil {
		t.Errorf("error casting to absolute path: %+v", err)
		return
	}

	bytes, err := os.ReadFile(path)
	if err != nil {
		t.Errorf("error reading file '%s': %+v", path, err)
		return
	}

	tokens, err := Lex(string(bytes))
	if err != nil {
		t.Errorf("error lexing file '%s': %+v", path, err)
	}

	ast, err := Parse(tokens)
	if err != nil {
		t.Errorf("error lexing file '%s': %+v", path, err)
	}

	marsh, _ := json.MarshalIndent(ast, "", "  ")
	t.Log(string(marsh))
}

func BenchmarkParser(b *testing.B) {
	path, err := filepath.Abs(fmt.Sprintf(
		"%s/test.ape",
		TEST_DIR,
	))
	if err != nil {
		b.Errorf("error casting to absolute path: %+v", err)
		return
	}

	bytes, err := os.ReadFile(path)
	if err != nil {
		b.Errorf("error reading file '%s': %+v", path, err)
		return
	}

	tokens, err := Lex(string(bytes))
	if err != nil {
		b.Errorf("error lexing file '%s': %+v", path, err)
	}

	for b.Loop() {
		_, err = Parse(tokens)
		if err != nil {
			b.Errorf("error lexing file '%s': %+v", path, err)
		}
	}
}
