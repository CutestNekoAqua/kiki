package command

import (
	"gitea.code-infection.com/efertone/kiki/pkg/feed"
	"github.com/spf13/cobra"
)

// AddFeedCmd command
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

func init() {
	AddFeedCmd.Flags().String("name", "", "Feed name (required)")
	AddFeedCmd.MarkFlagRequired("name")

	AddFeedCmd.Flags().String("user", "", "Name of the user (required)")
	AddFeedCmd.MarkFlagRequired("user")

	AddFeedCmd.Flags().String("url", "", "Feed URL (required)")
	AddFeedCmd.MarkFlagRequired("url")
}
