package unmarshal

import (
	"encoding/json"
	"testing"
)

const FILEPATH string = "../openapi.yaml"

func TestUnmarshal(t *testing.T) {
	comps, err := Unmarshal(FILEPATH)
	if err != nil {
		t.Errorf("error unmarshalling: %+v", err)
	}

	marsh, _ := json.MarshalIndent(comps, "", "  ")
	t.Log(string(marsh))
}
