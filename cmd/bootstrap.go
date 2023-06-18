package cmd

import "github.com/spf13/cobra"

var bootstrapCmd = &cobra.Command{
	Use:   "bootstrap",
	Short: "bootstrap will set up the control tower",
	Long: `bootstrap will set up the control tower, which is composed of a 
temporal cluster deployed in a cloud`,
	Run: func(cmd *cobra.Command, args []string) {
		panic("not implemented yet")
	},
}
