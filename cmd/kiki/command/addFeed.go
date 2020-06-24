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
		Short: "Add Feed to an Account",
		Run: func(cmd *cobra.Command, args []string) {
			name, _ := cmd.Flags().GetString("name")
			user, _ := cmd.Flags().GetString("user")
			url, _ := cmd.Flags().GetString("url")
			provider, _ := cmd.Flags().GetString("provider")

			err := feed.Add(name, user, url, provider)
			if err != nil {
				log.Fatalln(err)
			}
		},
	}

	requiredFlag(cmd, "name", "Feed name")
	requiredFlag(cmd, "user", "Name of the Account")
	requiredFlag(cmd, "url", "Feed URL")
	requiredFlag(cmd, "provider", "Provider")

	return cmd
}
