package fix

import (
	"github.com/spf13/cobra"

	"github.com/forbole/juno/v2/cmd/parse"

	fixblocks "github.com/forbole/juno/v2/cmd/fix/blocks"
)

// NewFixCmd returns the Cobra command allowing to fix some BDJuno bugs without having to re-sync the whole database
func NewFixCmd(parseCfg *parse.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:               "fix",
		Short:             "Apply some fixes without the need to re-syncing the whole database from scratch",
		PersistentPreRunE: runPersistentPreRuns(parse.ReadConfig(parseCfg)),
	}

	cmd.AddCommand(
		fixblocks.NewBlocksCmd(parseCfg),
	)

	return cmd
}

func runPersistentPreRuns(preRun func(_ *cobra.Command, _ []string) error) func(_ *cobra.Command, _ []string) error {
	return func(cmd *cobra.Command, args []string) error {
		if root := cmd.Root(); root != nil {
			if root.PersistentPreRunE != nil {
				err := root.PersistentPreRunE(root, args)
				if err != nil {
					return err
				}
			}
		}

		return preRun(cmd, args)
	}
}
