package cases

import (
	"context"
	"fmt"

	"github.com/h-fam/brain/loader/cases"
	"github.com/h-fam/brain/loader/cloud"

	"github.com/spf13/cobra"
)

func Add(ctx context.Context, parent *cobra.Command) {
	listCmd.RunE = cloud.ContextAction(ctx, list)
	addCmd.RunE = cloud.ContextAction(ctx, add)
	cmd.AddCommand(listCmd)
	cmd.AddCommand(addCmd)
	parent.AddCommand(cmd)
}

var (
	cmd = &cobra.Command{
		Use:   "case",
		Short: "case subcommands",
	}
	listCmd = &cobra.Command{
		Use:   "list",
		Short: "List case",
	}
	addCmd = &cobra.Command{
		Use:   "add",
		Short: "Add case",
	}
)

func list(ctx context.Context, cmd *cobra.Command, args []string) error {
	fmt.Println("list")
	return nil
}

func add(ctx context.Context, cmd *cobra.Command, args []string) error {
	var err error
	c := &cases.Case{}
	c.Manufacturer, err = cmd.Flags().GetString("manufacturer")
	c.Caliber, err = cmd.Flags().GetString("caliber")
	c.Primer, err = cmd.Flags().GetString("primer")

	cmd.Printf("Adding case: %+v\n", c)
	err = cases.Add(ctx, cloud.GetDSClient(ctx), c)
	if err != nil {
		return err
	}
	return nil
}

func init() {
	addCmd.Flags().String("manufacturer", "", "")
	addCmd.Flags().String("caliber", "", "")
	addCmd.Flags().String("primer", "", "")
}
