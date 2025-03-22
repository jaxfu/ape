package typescript

import (
	"github.com/jaxfu/ape/components"
)

func Typescript(components.AllComponents) ([]byte, error) {
	content := `
type ApeObject = {
	name: string;
	category?: string;	
	description?: string;
}
`

	// filehandler.Default.Write([]byte(content), "/Users/fraterhqc/repos/projects/ape/TEST.ts")
	return []byte(content), nil
	// bytes, err := json.Marshal(components)
	// if err != nil {
	// 	return []byte{}, fmt.Errorf("error marshalling: %+v", err)
	// }
	//
	// return bytes, nil
}
