package command

func init() {
	addAccountFlags()
	addFeedFlags()
}

func addAccountFlags() {
	AddAccountCmd.Flags().String("name", "", "Account name (required)")
	AddAccountCmd.MarkFlagRequired("name")

	AddAccountCmd.Flags().String("api-token", "", "API Token (required)")
	AddAccountCmd.MarkFlagRequired("api-token")
}

func addFeedFlags() {
	AddFeedCmd.Flags().String("name", "", "Feed name (required)")
	AddFeedCmd.MarkFlagRequired("name")

	AddFeedCmd.Flags().String("user", "", "Name of the user (required)")
	AddFeedCmd.MarkFlagRequired("user")

	AddFeedCmd.Flags().String("url", "", "Feed URL (required)")
	AddFeedCmd.MarkFlagRequired("url")
}
