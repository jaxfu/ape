package dev

import (
	"encoding/json"
	"fmt"
)

func PrettyPrint(a any) {
	m, err := json.MarshalIndent(a, "", "\t")
	if err != nil {
		fmt.Printf("PrettyPrint: %+v", err)
	}
	fmt.Printf("%+v\n", string(m))
}
