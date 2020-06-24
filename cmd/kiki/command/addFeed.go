package command

import (
	"log"

	"gitea.code-infection.com/efertone/kiki/pkg/model/feed"
	"gitea.code-infection.com/efertone/kiki/pkg/provider"
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
			p, _ := cmd.Flags().GetString("provider")

			if p == "xml" {
				var err error

				content, err := provider.Download(url)
				if err != nil {
					log.Printf("Error: %s\n", err)
					return
				}

				p, err = provider.XMLFeedTypeOf(content)
				if err != nil {
					log.Printf("Error: %s\n", err)
					return
				}

				log.Printf("Match found: xml -> %s", p)
			}

			if provider.NewProviderByName(p) == nil {
				log.Printf("Unknown provider: %s\n", p)
			}

			err := feed.Add(name, user, url, p)
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
