package command

import (
	"fmt"

	"gitea.code-infection.com/efertone/kiki/pkg/version"
	"github.com/spf13/cobra"
)

// Version command.
func Version() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "version",
		Short: "Print the version number of Kiki",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Fprintf(cmd.OutOrStdout(), "%s %s\n", version.AppName, version.Build)
		},
	}

	return cmd
}
