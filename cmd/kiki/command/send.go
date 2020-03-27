package command

import (
	"fmt"

	"gitea.code-infection.com/efertone/kiki/pkg/account"
	"gitea.code-infection.com/efertone/kiki/pkg/feed"
	"github.com/spf13/cobra"
)

var SendCmd = &cobra.Command{
	Use:   "send",
	Short: "Send new entries to Misskey",
	Run: func(cmd *cobra.Command, args []string) {
		for _, acc := range account.All() {
			for _, f := range feed.AllFor(acc) {
				next := feed.NextPendingEntries(f)
				fmt.Println(next)
			}
		}
	},
}
