package main

import (
	"encoding/json"
	"fmt"
)

type A struct {
	IVal int64
}

func main() {
	a := A{10}
	d, err := json.MarshalIndent(&a, "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(d))

	mm := make(map[string]interface{})

	mm["Hello"] = "World"
	mm["A"] = 1

	d, err = json.MarshalIndent(mm, "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(d))

}
