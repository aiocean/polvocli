package cmd

import "github.com/spf13/cobra"

var packageCmd = &cobra.Command{
	Use:     "package",
	Version: "0.0.1",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func init() {
	packageCmd.PersistentFlags().String("polvo_address", "127.0.0.1:8080", "Entry point url")
	packageCmd.AddCommand(packageCreateCmd)
	packageCmd.AddCommand(packageDeleteCmd)
	packageCmd.AddCommand(packageListCmd)
	packageCmd.AddCommand(packageGet)
}
