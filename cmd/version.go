package cmd

import "github.com/spf13/cobra"

var versionCmd = &cobra.Command{
	Use:     "version",
	Version: "0.0.1",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func init() {
	versionCmd.PersistentFlags().String("polvo_address", "127.0.0.1:8080", "Entry point url")
	versionCmd.AddCommand(versionUpdateCmd)
	versionCmd.AddCommand(versionListCmd)
	versionCmd.AddCommand(versionCreateCmd)
	versionCmd.AddCommand(versionDeleteCmd)
	versionCmd.AddCommand(versionGetCmd)
}
