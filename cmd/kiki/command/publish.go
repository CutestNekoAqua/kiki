package command

import (
	"fmt"
	"log"

	"gitea.code-infection.com/efertone/kiki/pkg/model/account"
	"gitea.code-infection.com/efertone/kiki/pkg/model/entry"
	"gitea.code-infection.com/efertone/kiki/pkg/model/feed"
	"gitea.code-infection.com/efertone/kiki/pkg/publisher"
	"github.com/spf13/cobra"
)

// Publish command.
func Publish() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "publish",
		Short: "Publish new entries to Misskey",
		Run: func(cmd *cobra.Command, args []string) {
			for _, acc := range account.All() {
				prov := publisher.NewTokenPublisherByName(acc.Publisher, acc.BaseURL, acc.APIToken)
				if prov == nil {
					log.Fatalf("Unknown Publisher: %s\n", acc.Publisher)
					continue
				}

				for _, f := range feed.AllFor(acc) {
					next := entry.NextPending(f)
					if next == nil {
						continue
					}
					err := prov.Publish(fmt.Sprintf("**%s**\n\n%s\n\n%s", next.Title, next.Excerpt(), next.Link))
					if err == nil {
						entry.MarkAsPosted(next)
						continue
					}
					log.Printf("[%s] %s\n", prov.Name(), err.Error())
				}
			}
		},
	}

	return cmd
}
