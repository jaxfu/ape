package internal

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

const (
	FILENAME    string = "TEST_FMT.ape"
	TEST_INDENT bool   = true
)

func TestFormatter(t *testing.T) {
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

	formatted, err := format(ast, TEST_INDENT)
	if err != nil {
		t.Errorf("error formatting file '%s': %+v", path, err)
	}

	outpath, _ := filepath.Abs(FILENAME)
	file, err := os.Create(outpath)
	if err != nil {
		t.Errorf("error creating file '%s': %+v", outpath, err)
	}
	defer file.Close()

	file.Write(formatted)
}

func BenchmarkFormatter(b *testing.B) {
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
		_, err = format(ast, TEST_INDENT)
		if err != nil {
			b.Errorf("error formatting file '%s': %+v", path, err)
		}
	}
}
