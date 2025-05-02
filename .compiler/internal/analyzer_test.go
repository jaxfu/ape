package internal

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestAnalyzer(t *testing.T) {
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

	parser := NewParser()
	ast, err := parser.Parse(tokens)
	if err != nil {
		t.Errorf("error parsing file '%s': %+v", path, err)
	}

	if err := Analyze(ast); err != nil {
		t.Errorf("error analyzing file '%s': %+v", path, err)
	}
}

func BenchmarkAnalyzer(b *testing.B) {
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

	parser := NewParser()
	ast, err := parser.Parse(tokens)
	if err != nil {
		b.Errorf("error lexing file '%s': %+v", path, err)
	}

	for b.Loop() {
		if err := Analyze(ast); err != nil {
			b.Errorf("error analyzing file '%s': %+v", path, err)
		}
	}
}
