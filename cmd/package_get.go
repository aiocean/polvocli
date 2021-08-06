package cmd

import (
	"github.com/spf13/cobra"
)

var packageGet = &cobra.Command{
	Use:     "get",
	Version: "0.0.1",
	RunE: packageGetHandler,
}

func packageGetHandler (cmd *cobra.Command, args []string) error {

	return nil
}

func init() {
	packageGet.PersistentFlags().String("orn", "", "package orn")
	packageGet.MarkPersistentFlagRequired("orn")
}
