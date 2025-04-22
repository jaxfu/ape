package internal

import (
	"encoding/json"
	"fmt"
)

func DeepCopy[T any](src *T) (T, error) {
	dest := new(T)

	marshalled, err := json.Marshal(*src)
	if err != nil {
		return *dest, fmt.Errorf("json.Marshal: %+v", err)
	}

	if err := json.Unmarshal(marshalled, dest); err != nil {
		return *dest, fmt.Errorf("json.Unmarshal: %+v", err)
	}

	return *dest, nil
}
