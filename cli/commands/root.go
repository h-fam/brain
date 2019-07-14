// Package commands provides a root level command for the cli tree.
package commands

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "root",
}

// Root returns the current Root command.
func Root() *cobra.Command {
	return rootCmd
}
