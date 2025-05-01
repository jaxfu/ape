package internal

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

const TEST_DIR string = "../test"

func TestLexer(t *testing.T) {
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
	marsh, _ := json.MarshalIndent(tokens, "", " ")
	t.Log(string(marsh))
}

func BenchmarkLexer(b *testing.B) {
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

	for b.Loop() {
		_, err = Lex(string(bytes))
		if err != nil {
			b.Errorf("error lexing file '%s': %+v", path, err)
		}
	}
}
