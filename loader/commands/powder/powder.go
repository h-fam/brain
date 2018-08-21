package powder

import (
	"context"
	"fmt"

	"github.com/marcushines/brain/loader/powder"

	"github.com/marcushines/brain/loader/cloud"
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
		Use:   "powder",
		Short: "powder subcommands",
	}
	listCmd = &cobra.Command{
		Use:   "list",
		Short: "List powder",
	}
	addCmd = &cobra.Command{
		Use:   "add",
		Short: "Add powder",
	}
)

func list(ctx context.Context, cmd *cobra.Command, args []string) error {
	fmt.Println("list")
	return nil
}

func add(ctx context.Context, cmd *cobra.Command, args []string) error {
	var err error
	p := &powder.Powder{}
	p.Manufacturer, err = cmd.Flags().GetString("manufacturer")
	p.Name, err = cmd.Flags().GetString("name")
	p.URL, err = cmd.Flags().GetString("url")

	cmd.Printf("Adding powder: %+v\n", p)
	err = powder.Add(ctx, cloud.GetDSClient(ctx), p)
	if err != nil {
		return err
	}
	return nil
}

func init() {
	addCmd.Flags().String("manufacturer", "", "")
	addCmd.Flags().String("name", "", "")
	addCmd.Flags().String("url", "", "")
}
