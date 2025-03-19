package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jaxfu/ape/pkg/parser"
	"github.com/jaxfu/ape/pkg/sql"
)

// 1: parse json
//		- Objects
//		- Routes
//		- Actions

func main() {
	// bytes, err := os.ReadFile("../Todo.json")
	// if err != nil {
	// 	log.Fatalf("ReadFile: %+v\n", err)
	// }
	//
	// parser := parser.NewParser()
	// rawJSON, err := parser.ParseJSON(bytes)
	// if err != nil {
	// 	fmt.Println("ParseJSON:", err)
	// 	return
	// }
	// fmt.Printf("JSON: %+v\n", rawJSON)

	parser := parser.NewParser()
	bytes, err := os.ReadFile("../Todo.toml")
	if err != nil {
		log.Fatalf("ReadFile: %+v\n", err)
	}

	rawTOML, err := parser.ParseTOML(bytes)
	if err != nil {
		fmt.Println("ParseTOML:", err)
		return
	}
	fmt.Printf("TOML: %+v\n", rawTOML)

	obj, err := parser.GenerateObject(rawTOML)
	if err != nil {
		fmt.Println("GenerateObject:", err)
		return
	}
	fmt.Printf("obj: %+v\n", obj)

	sqlHandler := sql.NewSqlHandler()
	sqlStr, err := sqlHandler.GenerateCreateTable(obj)
	if err != nil {
		fmt.Println("GenerateCreateTable:", err)
		return
	}
	fmt.Printf("%s\n", sqlStr)
}
