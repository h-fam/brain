package main

import (
	"fmt"

	"gopkg.in/yaml.v2"
	"source.cloud.google.com/hines-alloc/brain/cli/commands"
)

func main() {
	foo := map[string]string{
		"this": "that",
		"that": "1",
	}
	y, err := yaml.Marshal(foo)
	fmt.Println(string(y), err)
	cmd := commands.Root()
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
