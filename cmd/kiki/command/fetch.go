package command

import (
	"log"

	"gitea.code-infection.com/efertone/kiki/pkg/model/feed"
	"github.com/spf13/cobra"
)

// Fetch command.
func Fetch() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "fetch",
		Short: "Fetch all Feed contents",
		Run: func(cmd *cobra.Command, args []string) {
			for _, f := range feed.All() {
				log.Println(f)
				// feed.Get(f)
			}
		},
	}

	return cmd
}

/*
// Get fetches one Feed and store results in the database.
func Get(feed *model.Feed) {
		entries, err := Download(feed)
		if err != nil {
			log.Printf("[%s] Error: %s\n", feed.Name, err)
			return
		}

		db := database.NewDatabase()
		defer db.Close()

		for i := len(entries) - 1; i >= 0; i-- {
			var (
				entry = entries[i]
				count int
			)

			db.Connection().
				Model(&model.Entry{}).
				Where(&model.Entry{FeedID: entry.FeedID, EntryID: entry.EntryID}).
				Or(&model.Entry{FeedID: entry.FeedID, Link: entry.Link}).
				Count(&count)

			if count > 0 {
				continue
			}

			db.Connection().Create(entry)
		}
	}
*/
