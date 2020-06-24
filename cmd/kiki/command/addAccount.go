package command

import (
	"log"

	"gitea.code-infection.com/efertone/kiki/pkg/model/account"
	"github.com/spf13/cobra"
)

// AddAccount command.
func AddAccount() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "add-account",
		Short: "Add a new Account",
		Run: func(cmd *cobra.Command, args []string) {
			name, _ := cmd.Flags().GetString("name")
			token, _ := cmd.Flags().GetString("api-token")
			publisher, _ := cmd.Flags().GetString("publisher")
			baseURL, _ := cmd.Flags().GetString("base-url")

			err := account.Add(name, token, publisher, baseURL)
			if err != nil {
				log.Fatalln(err)
			}
		},
	}

	requiredFlag(cmd, "name", "Feed name")
	requiredFlag(cmd, "publisher", "Publisher")
	requiredFlag(cmd, "base-url", "Base URL like 'https://slippy.zyx'")
	requiredFlag(cmd, "api-token", "Account api-token")

	return cmd
}
