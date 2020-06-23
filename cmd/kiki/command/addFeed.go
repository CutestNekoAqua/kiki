package command

import (
	"log"

	"gitea.code-infection.com/efertone/kiki/pkg/model/feed"
	"github.com/spf13/cobra"
)

// AddFeed command.
func AddFeed() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "add-feed",
		Short: "Add Feed to a Misskey user",
		Run: func(cmd *cobra.Command, args []string) {
			name, _ := cmd.Flags().GetString("name")
			user, _ := cmd.Flags().GetString("user")
			url, _ := cmd.Flags().GetString("url")
			provider, _ := cmd.Flags().GetString("provider")
			feed.Add(name, user, url, provider)
		},
	}

	cmd.Flags().String("name", "", "Feed name (required)")

	if err := cmd.MarkFlagRequired("name"); err != nil {
		log.Fatalf("Lethal damage: %s\n", err)
	}

	cmd.Flags().String("user", "", "Name of the user (required)")

	if err := cmd.MarkFlagRequired("user"); err != nil {
		log.Fatalf("Lethal damage: %s\n", err)
	}

	cmd.Flags().String("url", "", "Feed URL (required)")

	if err := cmd.MarkFlagRequired("url"); err != nil {
		log.Fatalf("Lethal damage: %s\n", err)
	}

	cmd.Flags().String("provider", "", "Provider (required)")

	if err := cmd.MarkFlagRequired("provider"); err != nil {
		log.Fatalf("Lethal damage: %s\n", err)
	}

	return cmd
}
