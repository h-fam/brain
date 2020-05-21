package primer

import (
	"context"
	"fmt"

	"github.com/h-fam/brain/loader/cloud"
	"github.com/h-fam/brain/loader/primer"

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
		Use:   "primer",
		Short: "powder subcommands",
	}
	listCmd = &cobra.Command{
		Use:   "list",
		Short: "List primer",
	}
	addCmd = &cobra.Command{
		Use:   "add",
		Short: "Add primer",
	}
)

func list(ctx context.Context, cmd *cobra.Command, args []string) error {
	fmt.Println("list")
	return nil
}

func add(ctx context.Context, cmd *cobra.Command, args []string) error {
	var err error
	p := &primer.Primer{}
	p.Manufacturer, err = cmd.Flags().GetString("manufacturer")
	p.Name, err = cmd.Flags().GetString("name")
	p.Size, err = cmd.Flags().GetString("size")
	p.Type, err = cmd.Flags().GetString("type")
	p.URL, err = cmd.Flags().GetString("url")

	cmd.Printf("Adding primer: %+v\n", p)
	err = primer.Add(ctx, cloud.GetDSClient(ctx), p)
	if err != nil {
		return err
	}
	return nil
}

func init() {
	addCmd.Flags().String("manufacturer", "", "")
	addCmd.Flags().String("name", "", "")
	addCmd.Flags().String("size", "", "")
	addCmd.Flags().String("type", "", "")
	addCmd.Flags().String("url", "", "")
}
