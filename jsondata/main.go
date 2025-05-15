package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	data := generateJson()
	degenerateJson(data)
}
func generateJson() []uint8 {
	name := map[string]interface{}{
		"abc": "xyz",
		"atv": 55,
	}

	out, err := json.MarshalIndent(name, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", out)
	return out
}
func degenerateJson(str []uint8) {

	isvalid := json.Valid(str)
	if isvalid {
		var name map[string]interface{}
		json.Unmarshal(str, &name)
		fmt.Printf("%s\n", name)

	}

}
