package dev

import (
	"encoding/json"
	"fmt"
)

const (
	INDENT string = "  "
)

func PrettyPrint(a any) {
	m, err := json.MarshalIndent(
		a,
		"",
		INDENT,
	)
	if err != nil {
		fmt.Printf("PrettyPrint: %+v", err)
	}
	fmt.Printf("%+v\n", string(m))
}
