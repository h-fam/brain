package main

import (
	"fmt"

	"github.com/h-fam/brain/cli/commands"

	yaml "gopkg.in/yaml.v2"
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
