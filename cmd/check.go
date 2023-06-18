package cmd

import "github.com/spf13/cobra"

var checkCredentialsCmd = &cobra.Command{
	Use:   "check-credentials",
	Short: "check-credentials will check your cloud credentials",
	Long: `check-credentials will check your cloud credentials to ensure that
the credentials are valid and have sufficient permissions`,
	Run: func(cmd *cobra.Command, args []string) {
		panic("not implemented yet")
	},
}
