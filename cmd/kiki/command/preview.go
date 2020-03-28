package command

import (
	"fmt"
	"log"

	"gitea.code-infection.com/efertone/kiki/pkg/database"
	"gitea.code-infection.com/efertone/kiki/pkg/feed"
	"gitea.code-infection.com/efertone/kiki/pkg/model"
	"github.com/spf13/cobra"
)

var PreviewCmd = &cobra.Command{
	Use:   "preview",
	Short: "Preview desired content",
}

func init() {
	PreviewCmd.AddCommand(PreviewFetchCmd)

	PreviewFetchCmd.Flags().String("url", "", "Feed URL (required)")
	PreviewFetchCmd.MarkFlagRequired("url")
}

var PreviewFetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "Preview fetch",
	Run: func(cmd *cobra.Command, args []string) {
		url, _ := cmd.Flags().GetString("url")
		f := &model.Feed{URL: url, Name: "Preview"}

		db := database.NewDatabase()
		defer db.Close()
		db.Connection().NewRecord(f)

		entries, err := feed.Download(f)
		if err != nil {
			log.Printf("[%s] Error: %s\n", f.Name, err)
			return
		}
		for i := len(entries) - 1; i >= 0; i-- {
			entry := entries[i]
			fmt.Printf("== %s ==\n\n%s\n\n%s\n", entry.Title, entry.Excerpt(), entry.Link)
		}
	},
}
