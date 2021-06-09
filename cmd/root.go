package cmd

import "github.com/spf13/cobra"

var RootCmd = &cobra.Command{
	Use:     "polvocli",
	Short:   "Một công cụ hỗt rợ làm việc với các polvo.",
	Version: "0.0.1",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Usage()
	},
}

func init() {
	RootCmd.AddCommand(versionCmd)
}
