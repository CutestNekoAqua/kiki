package command

import (
	"gitea.code-infection.com/efertone/kiki/pkg/feed"
	"github.com/spf13/cobra"
)

var FetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "Fetch all Feed contents",
	Run: func(cmd *cobra.Command, args []string) {
		for _, f := range feed.All() {
			feed.Get(f)
		}
	},
}
