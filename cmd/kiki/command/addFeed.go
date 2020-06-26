package command

import (
	"fmt"

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
					fmt.Fprintf(cmd.OutOrStderr(), "Error: %s\n", err)
					return
				}

				p, err = provider.XMLFeedTypeOf(content)
				if err != nil {
					fmt.Fprintf(cmd.OutOrStderr(), "Error: %s\n", err)
					return
				}

				fmt.Fprintf(cmd.OutOrStdout(), "Match found: xml -> %s\n", p)
			}

			if provider.NewProviderByName(p) == nil {
				fmt.Fprintf(cmd.OutOrStderr(), "Unknown provider: %s\n", p)
			}

			err := feed.Add(name, user, url, p)
			if err != nil {
				fmt.Fprintln(cmd.OutOrStderr(), err)
			}
		},
	}

	requiredFlag(cmd, "name", "Feed name")
	requiredFlag(cmd, "user", "Name of the Account")
	requiredFlag(cmd, "url", "Feed URL")
	requiredFlag(cmd, "provider", "Provider")

	return cmd
}
