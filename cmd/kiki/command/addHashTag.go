package command

import (
	"fmt"

	"gitea.code-infection.com/efertone/kiki/pkg/model/feed"
	"gitea.code-infection.com/efertone/kiki/pkg/model/hashtag"
	"github.com/spf13/cobra"
)

// AddHashTag command.
func AddHashTag() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "add-hashtag",
		Short: "Add HashTag to a Feed",
		Run: func(cmd *cobra.Command, args []string) {
			feedID, _ := cmd.Flags().GetUint("feed-id")
			value, _ := cmd.Flags().GetString("value")

			_, err := feed.FindByID(feedID)
			if err != nil {
				fmt.Fprintln(cmd.OutOrStderr(), err)
				return
			}

			err = hashtag.Add(feedID, value)
			if err != nil {
				fmt.Fprintln(cmd.OutOrStderr(), err)
			}
		},
	}

	requiredUintFlag(cmd, "feed-id", "ID of the Feed")
	requiredFlag(cmd, "value", "HashTag value")

	return cmd
}
