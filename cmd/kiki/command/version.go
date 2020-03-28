package command

import (
	"fmt"

	"gitea.code-infection.com/efertone/kiki/pkg/version"
	"github.com/spf13/cobra"
)

var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Kiki",
	Run: func(cmd *cobra.Command, args []string) {
		if version.Tag != "" {
			fmt.Printf("%s %s\n", version.AppName, version.Tag)
			return
		}
		fmt.Printf("%s devel-%s\n", version.AppName, version.Build)
	},
}