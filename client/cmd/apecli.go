package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/jaxfu/ape/components"
)

const (
	PORT int = 5173
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
	_, err = http.Post(
		"http://localhost:5000/api/components",
		"application/json",
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		log.Fatalf("%+v\n", err)
	}

	router := http.NewServeMux()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fp, _ := filepath.Abs("web/vanilla/index.html")
		http.ServeFile(w, r, fp)
	})
	router.HandleFunc("/index.js", func(w http.ResponseWriter, r *http.Request) {
		fp, _ := filepath.Abs("web/vanilla/index.js")
		http.ServeFile(w, r, fp)
	})
	router.HandleFunc("/styles.css", func(w http.ResponseWriter, r *http.Request) {
		fp, _ := filepath.Abs("web/vanilla/styles.css")
		http.ServeFile(w, r, fp)
	})

	server := http.Server{
		Addr:    fmt.Sprintf("localhost:%d", PORT),
		Handler: router,
	}

	fmt.Printf("Listening on port %d...", PORT)
	log.Fatal(server.ListenAndServe())

	// object := components.Object{
	// 	ComponentMetadata: components.ComponentMetadata{
	// 		ComponentType: "OBJECT",
	// 		ComponentId:   "TEST",
	// 		Name:          "TEST",
	// 		IsRoot:        true,
	// 	},
	// 	Props: components.PropsMap{
	// 		"username": components.Prop{
	// 			ComponentMetadata: components.ComponentMetadata{
	// 				ComponentType: "PROP",
	// 				ComponentId:   "",
	// 				Name:          "username",
	// 				IsRoot:        false,
	// 			},
	// 			PropMetadata: components.PropMetadata{
	// 				PropType: "INT",
	// 				IsArray:  false,
	// 			},
	// 			Constraints: nil,
	// 		},
	// 	},
	// }
	//
	// jsonData, err := json.Marshal(object)
	// if err != nil {
	// 	panic(err)
	// }
	//
	// // Create a POST request
	// _, err = http.Post(
	// 	"http://localhost:5000/api/components",
	// 	"application/json",
	// 	bytes.NewBuffer(jsonData),
	// )
	// if err != nil {
	// 	log.Fatalf("%+v\n", err)
	// }
	// defer resp.Body.Close()
	// body, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Fatalf("%+v\n", err)
	// }
	// fmt.Printf("POST: %s\n", resp.Status)
	// fmt.Println(string(body))
	//
	// // Get Components
	// resp, err = http.Get(
	// 	"http://localhost:5000/api/components",
	// )
	// for k, v := range resp.Header {
	// 	fmt.Printf("%s: %+v\n", k, v)
	// }
	// if err != nil {
	// 	log.Fatalf("%+v\n", err)
	// }
	// defer resp.Body.Close()
	// body, err = io.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Fatalf("%+v\n", err)
	// }
	// fmt.Printf("GET: %s\n", resp.Status)
	// fmt.Printf("%s\n", string(body))

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
