package main

import (
	"log"

	"gitea.code-infection.com/efertone/kiki/cmd/kiki/command"
	"gitea.code-infection.com/efertone/kiki/pkg/database"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	configFile string
	rootCmd    = &cobra.Command{
		Use:   "kiki",
		Short: "RSS Delivery Service",
		Long:  "Deliver RSS items on Misskey",
	}
)

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().Bool("debug", false, "spam stdout with debug information")
	rootCmd.PersistentFlags().StringVar(
		&configFile,
		"config",
		"",
		"config file (default is $HOME/.config/kiki/config.yaml)")

	viper.BindPFlag("debug", rootCmd.PersistentFlags().Lookup("debug"))

	rootCmd.AddCommand(command.VersionCmd)
	rootCmd.AddCommand(command.AddAccountCmd)
	rootCmd.AddCommand(command.AddFeedCmd)
	rootCmd.AddCommand(command.FetchCmd)
	rootCmd.AddCommand(command.PreviewCmd)
	rootCmd.AddCommand(command.SendCmd)
}

func initConfig() {
	if configFile != "" {
		viper.SetConfigFile(configFile)
	} else {
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath("$HOME/.config/kiki")
		viper.AddConfigPath(".")
	}

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Fatal error config file: %s\n", err)
	}

	if viper.GetBool("debug") {
		log.Println("Using config file:", viper.ConfigFileUsed())
	}

	boot()
}

func boot() {
	database.Configure(&database.ConnectionDetails{
		User:     viper.GetString("database.user"),
		Password: viper.GetString("database.password"),
		Name:     viper.GetString("database.name"),
		Host:     viper.GetString("database.host"),
		Port:     viper.GetInt("database.port"),
	})
	database.NewDatabase().Migrate()
}

func main() {
	rootCmd.Execute()
}
