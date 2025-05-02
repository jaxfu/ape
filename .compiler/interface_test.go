package compiler

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

const (
	TEST_DIR = "test"
)

var compiler = NewCompiler()

func TestCompiler(t *testing.T) {
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
		t.Errorf("error reading file %s: %+v", path, err)
		return
	}

	_, err = compiler.File(path, bytes)
	if err != nil {
		t.Errorf("error compiling file %s: %+v", path, err)
		return
	}
}

func BenchmarkCompiler(b *testing.B) {
	for b.Loop() {
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
			b.Errorf("error reading file %s: %+v", path, err)
			return
		}

		_, err = compiler.File(path, bytes)
		if err != nil {
			b.Errorf("error compiling file %s: %+v", path, err)
			return
		}

	}
}
