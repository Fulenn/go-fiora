package cmd

import (
	"fmt"
	"github.com/Fulenn/go-fiora/internal/builder"
	"github.com/spf13/cobra"
	"os"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   `fiora`,
	Short: ``,
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			builder.Runer(args[0]) // Passes the first argument as the project name
		} else {
			fmt.Println("Please provide a matchup")
			// Optionally, you can return an error or exit the command if no argument is provided
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
