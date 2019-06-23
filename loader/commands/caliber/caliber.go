package caliber

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	"source.cloud.google.com/hines-alloc/brain/loader/caliber"
	"source.cloud.google.com/hines-alloc/brain/loader/cloud"
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
		Use:   "caliber",
		Short: "caliber subcommands",
	}
	listCmd = &cobra.Command{
		Use:   "list",
		Short: "List caliber",
	}
	addCmd = &cobra.Command{
		Use:   "add",
		Short: "Add caliber",
	}
)

func list(ctx context.Context, cmd *cobra.Command, args []string) error {
	fmt.Println("list")
	return nil
}

func add(ctx context.Context, cmd *cobra.Command, args []string) error {
	var err error
	c := &caliber.Caliber{}
	c.Name, err = cmd.Flags().GetString("name")
	c.Diameter, err = cmd.Flags().GetInt64("diameter")
	c.URL, err = cmd.Flags().GetString("url")

	cmd.Printf("Adding caliber: %+v\n", c)
	err = caliber.Add(ctx, cloud.GetDSClient(ctx), c)
	if err != nil {
		return err
	}
	return nil
}

func init() {
	addCmd.Flags().String("name", "", "")
	addCmd.Flags().Int64("diameter", 0, "")
	addCmd.Flags().String("url", "", "")
}
