package cmd

import (
	"fmt"
	"io"

	"github.com/spf13/cobra"
	v1 "pkg.aiocean.dev/polvogo/aiocean/polvo/v1"
)

var packageDeleteCmd = &cobra.Command{
	Use:     "delete",
	Version: "0.0.1",
	RunE: packageDeleteHandler,
}

func packageDeleteHandler (cmd *cobra.Command, args []string) error {
	polvoAddress, err := cmd.Flags().GetString("polvo_address")
	if err != nil {
		return err
	}

	orn, err := cmd.Flags().GetString("orn")
	if err != nil {
		return err
	}

	client, err := newPolvoClient(cmd.Context(), polvoAddress)
	if err != nil {
		return err
	}

	request := &v1.DeletePackageRequest{
		Orn: orn,
	}

	stream, err := client.DeletePackage(cmd.Context(), request)
	if err != nil {
		return err
	}

	for {
		response, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		fmt.Println(response.GetMessage())
	}
}

func init() {
	packageDeleteCmd.Flags().String("orn", "", "package orn")
}
