package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "bravo",
	Short: "Bravo is a cli for generating projects templates",
	Long: `
A simple cli that eliminate the guess work in generating go apps.

Create your go apps by following the industry standard project layout,
with a wide range of customization from building APIs,Webapps,CLIs
with git integrations.

-> Focus on writing code and thinking of business logic!
<- Tool build for developers by developers.

A helpful documentation and next steps -> https://github.com/bravo-app/cli`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Bravo is a cli for generating projects templates")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
