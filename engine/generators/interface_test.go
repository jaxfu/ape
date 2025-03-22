package generators

import (
	"encoding/json"
	"log"
	"os"
	"testing"
)

type testInputs struct{}

var inputs testInputs

const (
	TEST_FILE string = "testdata.json"
)

func TestMain(m *testing.M) {
	bytes, err := os.ReadFile(TEST_FILE)
	if err != nil {
		log.Fatalf("error reading '%s': %+v\n", TEST_FILE, err)
	}

	if err := json.Unmarshal(bytes, &inputs); err != nil {
		os.Exit(1)
	}

	// code := m.Run()
	// os.Exit(code)
}
