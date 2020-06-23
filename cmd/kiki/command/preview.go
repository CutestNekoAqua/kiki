package command

import (
	"fmt"
	"log"

	"gitea.code-infection.com/efertone/kiki/pkg/provider"
	"github.com/spf13/cobra"
)

// Preview command.
func Preview() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "preview",
		Short: "Preview desired content",
	}

	cmd.AddCommand(PreviewFetch())

	return cmd
}

// PreviewFetch : Preview -> Fetch command.
func PreviewFetch() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "fetch",
		Short: "Preview fetch",
		Run: func(cmd *cobra.Command, args []string) {
			url, _ := cmd.Flags().GetString("url")
			p, _ := cmd.Flags().GetString("provider")

			content, err := provider.Download(url)

			if err != nil {
				log.Printf("[Preview] Error: %s\n", err)
			}

			handler := provider.NewProviderByName(p)
			if handler == nil {
				log.Printf("[Preview] Provider not found: %s\n", p)
			}

			entries := handler.Parse(content)
			for i := len(entries) - 1; i >= 0; i-- {
				entry := entries[i].ToModel()
				fmt.Printf("== %s ==\n\n%s\n\n%s\n", entry.Title, entry.Excerpt(), entry.Link)
			}
		},
	}

	cmd.Flags().String("url", "", "Feed URL (required)")

	if err := cmd.MarkFlagRequired("url"); err != nil {
		log.Fatalf("Lethal damage: %s\n", err)
	}

	cmd.Flags().String("provider", "", "Provider (required)")

	if err := cmd.MarkFlagRequired("provider"); err != nil {
		log.Fatalf("Lethal damage: %s\n", err)
	}

	return cmd
}
