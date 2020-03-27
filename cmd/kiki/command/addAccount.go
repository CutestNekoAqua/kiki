package command

import (
	"gitea.code-infection.com/efertone/kiki/pkg/account"
	"github.com/spf13/cobra"
)

var AddAccountCmd = &cobra.Command{
	Use:   "add-account",
	Short: "Add Misskey user",
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		token, _ := cmd.Flags().GetString("api-token")
		account.Add(name, token)
	},
}
