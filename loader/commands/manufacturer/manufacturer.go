package manufacturer

import (
	"context"
	"fmt"

	"github.com/h-fam/brain/loader/cloud"
	"github.com/h-fam/brain/loader/manufacturer"

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
		Use:   "manufacturer",
		Short: "manufacturer subcommands",
	}
	listCmd = &cobra.Command{
		Use:   "list",
		Short: "List manufacturer",
	}
	addCmd = &cobra.Command{
		Use:   "add",
		Short: "Add manufacturer",
	}
)

func list(ctx context.Context, cmd *cobra.Command, args []string) error {
	fmt.Println("list")
	return nil
}

func add(ctx context.Context, cmd *cobra.Command, args []string) error {
	var err error
	m := &manufacturer.Manufacturer{}
	m.Name, err = cmd.Flags().GetString("name")
	m.URL, err = cmd.Flags().GetString("url")

	cmd.Printf("Adding manufacturer: %+v\n", m)
	err = manufacturer.Add(ctx, cloud.GetDSClient(ctx), m)
	if err != nil {
		return err
	}
	return nil
}

func init() {
	addCmd.Flags().String("name", "", "")
	addCmd.Flags().String("url", "", "")
}
