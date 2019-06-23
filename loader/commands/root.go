package commands

import (
	"context"

	"github.com/spf13/cobra"
	"source.cloud.google.com/hines-alloc/brain/loader/commands/bullet"
	"source.cloud.google.com/hines-alloc/brain/loader/commands/caliber"
	"source.cloud.google.com/hines-alloc/brain/loader/commands/cases"
	"source.cloud.google.com/hines-alloc/brain/loader/commands/manufacturer"
	"source.cloud.google.com/hines-alloc/brain/loader/commands/powder"
	"source.cloud.google.com/hines-alloc/brain/loader/commands/primer"
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
