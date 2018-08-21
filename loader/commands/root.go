package commands

import (
	"context"

	"github.com/marcushines/brain/loader/commands/bullet"
	"github.com/marcushines/brain/loader/commands/caliber"
	"github.com/marcushines/brain/loader/commands/cases"
	"github.com/marcushines/brain/loader/commands/manufacturer"
	"github.com/marcushines/brain/loader/commands/powder"
	"github.com/marcushines/brain/loader/commands/primer"
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
