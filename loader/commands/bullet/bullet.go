package bullet

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	"source.cloud.google.com/hines-alloc/brain/loader/bullet"
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
		Use:   "bullet",
		Short: "Bullet subcommands",
	}
	listCmd = &cobra.Command{
		Use:   "list",
		Short: "List bullets",
	}
	addCmd = &cobra.Command{
		Use:   "add",
		Short: "Add bullets",
	}
)

func list(ctx context.Context, cmd *cobra.Command, args []string) error {
	fmt.Println("list")
	return nil
}

func add(ctx context.Context, cmd *cobra.Command, args []string) error {
	var err error
	b := &bullet.Bullet{}
	b.Manufacturer, err = cmd.Flags().GetString("manufacturer")
	b.Name, err = cmd.Flags().GetString("name")
	b.Caliber, err = cmd.Flags().GetString("caliber")
	b.Weight, err = cmd.Flags().GetInt("weight")
	b.Shape, err = cmd.Flags().GetString("shape")

	cmd.Println("Adding bullet: %s", b)
	err = bullet.Add(ctx, cloud.GetDSClient(ctx), b)
	if err != nil {
		return err
	}
	return nil
}

func init() {
	addCmd.Flags().String("manufacturer", "", "")
	addCmd.Flags().String("name", "", "")
	addCmd.Flags().String("caliber", "", "")
	addCmd.Flags().String("shape", "", "")
	addCmd.Flags().Int("weight", 0, "")
}
