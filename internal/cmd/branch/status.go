package branch

import (
	"context"
	"fmt"

	"github.com/planetscale/cli/internal/cmdutil"
	"github.com/planetscale/cli/internal/config"
	"github.com/planetscale/cli/internal/printer"

	"github.com/planetscale/planetscale-go/planetscale"

	"github.com/spf13/cobra"
)

// StatusCmd gets the status of a database branch using the PlanetScale API.
func StatusCmd(cfg *config.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "status <db_name> <branch_name>",
		Short: "Check the status of a branch of a database",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.Background()
			if len(args) != 2 {
				return cmd.Usage()
			}

			client, err := cfg.NewClientFromConfig()
			if err != nil {
				return err
			}

			source := args[0]
			branch := args[1]

			end := cmdutil.PrintProgress(fmt.Sprintf("Getting status for branch %s in %s...", cmdutil.BoldBlue(branch), cmdutil.BoldBlue(source)))
			defer end()
			status, err := client.DatabaseBranches.GetStatus(ctx, &planetscale.GetDatabaseBranchStatusRequest{
				Organization: cfg.Organization,
				Database:     source,
				Branch:       branch,
			})
			if err != nil {
				return err
			}

			isJSON, err := cmd.Flags().GetBool("json")
			if err != nil {
				return err
			}

			end()
			err = printer.PrintOutput(isJSON, printer.NewDatabaseBranchStatusPrinter(status))
			if err != nil {
				return err
			}

			return nil
		},
	}

	return cmd
}