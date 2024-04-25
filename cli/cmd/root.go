package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "uwuC2",
	Short: "uwuC2 is a lightweight C2 client",
	Long:  `A fast and flexible C2 client built with love by and for security researchers.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to uwuC2!")
	},
}

func Execute() error {
	return rootCmd.Execute()
}
