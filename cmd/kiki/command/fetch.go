package command

import (
	"log"

	"gitea.code-infection.com/efertone/kiki/pkg/database"
	"gitea.code-infection.com/efertone/kiki/pkg/model/entry"
	"gitea.code-infection.com/efertone/kiki/pkg/model/feed"
	"gitea.code-infection.com/efertone/kiki/pkg/provider"
	"github.com/spf13/cobra"
)

// Fetch command.
func Fetch() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "fetch",
		Short: "Fetch all Feed contents",
		Run: func(cmd *cobra.Command, args []string) {
			db := database.NewDatabase()
			defer db.Close()

			for _, f := range feed.All() {
				content, err := provider.Download(f.URL)
				if err != nil {
					log.Printf("[Fetch] Error: %s\n", err)
					continue
				}

				handler := provider.NewProviderByName(f.Provider)
				if handler == nil {
					log.Printf(
						"[Fetch] '%s' Provider not found for '%s'\n",
						f.Provider,
						f.Name,
					)
					continue
				}

				entries := handler.Parse(content)
				for i := len(entries) - 1; i >= 0; i-- {
					data := entries[i].ToModel()
					data.FeedID = f.ID

					if entry.Exists(data) {
						continue
					}

					db.Connection().Create(data)
				}
			}
		},
	}

	return cmd
}
