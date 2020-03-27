package command

import (
	"gitea.code-infection.com/efertone/kiki/pkg/account"
	"gitea.code-infection.com/efertone/kiki/pkg/feed"
	"gitea.code-infection.com/efertone/kiki/pkg/misskey"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var SendCmd = &cobra.Command{
	Use:   "send",
	Short: "Send new entries to Misskey",
	Run: func(cmd *cobra.Command, args []string) {
		for _, acc := range account.All() {
			mkClient := misskey.NewClient(viper.GetString("misskey.baseurl"), acc.APIToken)
			for _, f := range feed.AllFor(acc) {
				next := feed.NextPendingEntries(f)
				if next == nil {
					continue
				}
				if mkClient.CreateNote(next.Excerpt()) {
					feed.MarAsPosted(next)
				}
			}
		}
	},
}
