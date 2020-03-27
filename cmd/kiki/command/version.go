package command

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Kiki",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("version number goes here")
		fmt.Println(viper.GetString("database"))
	},
}
