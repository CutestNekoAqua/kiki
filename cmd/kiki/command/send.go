package command

import (
	"fmt"
	"log"

	"gitea.code-infection.com/efertone/kiki/pkg/account"
	"gitea.code-infection.com/efertone/kiki/pkg/feed"
	"gitea.code-infection.com/efertone/kiki/pkg/provider"
	"github.com/spf13/cobra"
)

// Send command.
func Send() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "send",
		Short: "Send new entries to Misskey",
		Run: func(cmd *cobra.Command, args []string) {
			for _, acc := range account.All() {
				prov := provider.NewTokenProviderByName(acc.Provider, acc.BaseURL, acc.APIToken)

				for _, f := range feed.AllFor(acc) {
					next := feed.NextPendingEntries(f)
					if next == nil {
						continue
					}
					err := prov.Pulish(fmt.Sprintf("**%s**\n\n%s\n\n%s", next.Title, next.Excerpt(), next.Link))
					if err == nil {
						feed.MarAsPosted(next)
						continue
					}
					log.Printf("[%s] %s\n", prov.Name(), err.Error())
				}
			}
		},
	}

	return cmd
}
