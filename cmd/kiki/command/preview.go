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
				fmt.Fprintf(cmd.OutOrStderr(), "[Preview] Error: %s\n", err)
				return
			}

			if p == "xml" {
				p, err = provider.XMLFeedTypeOf(content)
				if err != nil {
					fmt.Fprintf(cmd.OutOrStderr(), "[Preview] Error: %s\n", err)
					return
				}

				fmt.Fprintf(cmd.OutOrStdout(), "Match found: xml -> %s", p)
			}

			handler := provider.NewProviderByName(p)
			if handler == nil {
				fmt.Fprintf(cmd.OutOrStderr(), "[Preview] Provider not found: %s\n", p)
				return
			}

			entries := handler.Parse(content)
			for i := len(entries) - 1; i >= 0; i-- {
				entry := entries[i].ToModel()
				fmt.Fprintf(
					cmd.OutOrStdout(),
					"== %s ==\n\n%s\n\n%s\n",
					entry.Title,
					entry.Excerpt(),
					entry.Link,
				)
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
