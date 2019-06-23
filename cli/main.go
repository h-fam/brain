package main

import (
	"fmt"

	"gopkg.in/yaml.v2"
)

func main() {
	foo := map[string]string{
		"this": "that",
		"that": "1",
	}
	y, err := yaml.Marshal(foo)
	fmt.Println(string(y), err)
}
