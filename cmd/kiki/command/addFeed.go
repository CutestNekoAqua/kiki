package command

import (
	"gitea.code-infection.com/efertone/kiki/pkg/feed"
	"github.com/spf13/cobra"
)

var AddFeedCmd = &cobra.Command{
	Use:   "add-feed",
	Short: "Add Feed to a Misskey user",
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		user, _ := cmd.Flags().GetString("user")
		url, _ := cmd.Flags().GetString("url")
		feed.Add(name, user, url)
	},
}
