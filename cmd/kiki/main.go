package main

import (
	"log"

	"gitea.code-infection.com/efertone/kiki/cmd/kiki/command"
	"gitea.code-infection.com/efertone/kiki/pkg/database"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// RootCommand is the root kiki application.
func RootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "kiki",
		Short: "RSS Delivery Service",
		Long:  "Deliver RSS items on Misskey",
	}

	cmd.PersistentFlags().Bool("debug", false, "spam stdout with debug information")
	cmd.PersistentFlags().String(
		"config",
		"",
		"config file (default is $HOME/.config/kiki/config.yaml)")

	if err := viper.BindPFlag("debug", cmd.PersistentFlags().Lookup("debug")); err != nil {
		log.Fatalf("Lethal damage: %s\n", err)
	}

	if err := viper.BindPFlag("config", cmd.PersistentFlags().Lookup("config")); err != nil {
		log.Fatalf("Lethal damage: %s\n", err)
	}

	cmd.AddCommand(command.Version())
	cmd.AddCommand(command.AddAccount())
	cmd.AddCommand(command.AddFeed())
	cmd.AddCommand(command.AddHashTag())
	cmd.AddCommand(command.Fetch())
	cmd.AddCommand(command.Preview())
	cmd.AddCommand(command.Publish())
	cmd.AddCommand(command.ListFeeds())
	cmd.AddCommand(command.ListAccounts())

	cobra.OnInitialize(initConfig)

	return cmd
}

func initConfig() {
	if viper.GetString("config") != "" {
		viper.SetConfigFile(viper.GetString("config"))
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

	bootDatabase()
}

func bootDatabase() {
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
	root := RootCommand()

	_ = root.Execute()
}
