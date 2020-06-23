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
		Short: "Add Misskey user",
		Run: func(cmd *cobra.Command, args []string) {
			name, _ := cmd.Flags().GetString("name")
			token, _ := cmd.Flags().GetString("api-token")
			publisher, _ := cmd.Flags().GetString("publisher")
			baseURL, _ := cmd.Flags().GetString("base-url")
			account.Add(name, token, publisher, baseURL)
		},
	}

	cmd.Flags().String("name", "", "Account name (required)")

	if err := cmd.MarkFlagRequired("name"); err != nil {
		log.Fatalf("Lethal damage: %s\n", err)
	}

	cmd.Flags().String("publisher", "", "Publisher (required)")

	if err := cmd.MarkFlagRequired("publisher"); err != nil {
		log.Fatalf("Lethal damage: %s\n", err)
	}

	cmd.Flags().String("base-url", "", "base-url (required)")

	if err := cmd.MarkFlagRequired("base-url"); err != nil {
		log.Fatalf("Lethal damage: %s\n", err)
	}

	cmd.Flags().String("api-token", "", "Account api-token (required)")

	if err := cmd.MarkFlagRequired("api-token"); err != nil {
		log.Fatalf("Lethal damage: %s\n", err)
	}

	return cmd
}
