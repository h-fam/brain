package commands

import (
	"context"

	"hfam/brain/loader/commands/bullet"
	"hfam/brain/loader/commands/caliber"
	"hfam/brain/loader/commands/cases"
	"hfam/brain/loader/commands/manufacturer"
	"hfam/brain/loader/commands/powder"
	"hfam/brain/loader/commands/primer"

	"github.com/spf13/cobra"
)

var root = &cobra.Command{
	Use:   "loader",
	Short: "Loader is the CLI interface to the Loader API",
}

func AddRoot(ctx context.Context) {
	bullet.Add(ctx, root)
	caliber.Add(ctx, root)
	manufacturer.Add(ctx, root)
	powder.Add(ctx, root)
	primer.Add(ctx, root)
	cases.Add(ctx, root)
}

func Run() error {
	return root.Execute()
}
