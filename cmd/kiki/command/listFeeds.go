package command

import (
	"fmt"
	"strings"

	"gitea.code-infection.com/efertone/kiki/pkg/model/feed"
	"github.com/spf13/cobra"
)

// ListFeeds command.
func ListFeeds() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "list-feeds",
		Short: "List all Feeds",
		Run: func(cmd *cobra.Command, args []string) {
			feeds := feed.All()

			for _, f := range feeds {
				tags := feed.PrefixedHashTagValuesFor(f)

				fmt.Fprintf(
					cmd.OutOrStdout(),
					"%4d | %-20.20s | %-20.20s | %s\n",
					f.ID,
					f.User,
					f.Name,
					strings.Join(tags, " "),
				)
			}
		},
	}

	return cmd
}
