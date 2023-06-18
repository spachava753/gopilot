package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gopilot",
	Short: "gopilot is a tool to run ML workloads in a cost efficient manner",
	Long: `gopilot is a tool to run ML workloads in a cost efficient manner, 
utilizing preemptable instances across availability zones, regions and even clouds`,
}

func Execute() {
	// register the root's subcommands here
	rootCmd.AddCommand(bootstrapCmd)
	rootCmd.AddCommand(checkCredentialsCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
