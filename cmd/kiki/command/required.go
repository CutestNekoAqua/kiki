package command

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

func requiredFlag(cmd *cobra.Command, name, description string) {
	cmd.Flags().String(
		name,
		"",
		fmt.Sprintf("%s (required)", description),
	)

	if err := cmd.MarkFlagRequired(name); err != nil {
		log.Fatalf("Lethal damage: %s\n", err)
	}
}

func requiredUintFlag(cmd *cobra.Command, name, description string) {
	cmd.Flags().Uint(
		name,
		0,
		fmt.Sprintf("%s (required)", description),
	)

	if err := cmd.MarkFlagRequired(name); err != nil {
		log.Fatalf("Lethal damage: %s\n", err)
	}
}
