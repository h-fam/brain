package cloud

import (
	"context"

	"github.com/spf13/cobra"
)

type BaseRun func(*cobra.Command, []string) error

type ContextRun func(context.Context, *cobra.Command, []string) error

func ContextAction(ctx context.Context, f ContextRun) BaseRun {
	return func(cmd *cobra.Command, args []string) error {
		return f(ctx, cmd, args)
	}
}
