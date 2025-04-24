package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/jaxfu/ape/components"
)

func main() {
	object := components.Object{
		ComponentMetadata: components.ComponentMetadata{
			ComponentType: "OBJECT",
			ComponentId:   "TEST",
			Name:          "TEST",
			IsRoot:        true,
		},
		Props: components.PropsMap{
			"username": components.Prop{
				ComponentMetadata: components.ComponentMetadata{
					ComponentType: "PROP",
					ComponentId:   "",
					Name:          "username",
					IsRoot:        false,
				},
				PropMetadata: components.PropMetadata{
					PropType: "INT",
					IsArray:  false,
				},
				Constraints: nil,
			},
		},
	}

	jsonData, err := json.Marshal(object)
	if err != nil {
		panic(err)
	}

	// Create a POST request
	resp, err := http.Post(
		"http://localhost:5000/api/components",
		"application/json",
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		log.Fatalf("%+v\n", err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("%+v\n", err)
	}
	defer resp.Body.Close()
	fmt.Printf("status: %s\n", resp.Status)
	fmt.Println(string(body))

	// Read the response
	// var result map[string]any
	// json.NewDecoder(resp.Body).Decode(&result)
	// PrettyPrint(result)
}

func PrettyPrint(a any) {
	if a == nil {
		return
	}

	m, err := json.MarshalIndent(
		a,
		"",
		" ",
	)
	if err != nil {
		fmt.Printf("json.MarshalIndent: %+v", err)
	}

	fmt.Printf("%+v\n", string(m))
}
