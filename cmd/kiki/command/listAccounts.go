package command

import (
	"fmt"

	"gitea.code-infection.com/efertone/kiki/pkg/model/account"
	"gitea.code-infection.com/efertone/kiki/pkg/model/feed"
	"github.com/spf13/cobra"
)

// ListAccounts command.
func ListAccounts() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "list-accounts",
		Short: "List all accounts",
		Run: func(cmd *cobra.Command, args []string) {
			accounts := account.All()

			for _, acc := range accounts {
				feeds := feed.AllFor(acc)

				fmt.Fprintf(
					cmd.OutOrStdout(),
					"%4d | %-20.20s | %-10.10s | %-20.20s | %d feed(s)\n",
					acc.ID,
					acc.Name,
					acc.Publisher,
					acc.BaseURL,
					len(feeds),
				)
			}
		},
	}

	return cmd
}
