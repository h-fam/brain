package main

import (
	"context"
	"fmt"
	"os"

	"source.cloud.google.com/hines-alloc/brain/loader/commands"
)

func main() {
	commands.AddRoot(context.Background())
	if err := commands.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}