package cmd

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/spf13/cobra"
	v1 "pkg.aiocean.dev/polvogo/aiocean/polvo/v1"
)

var versionListCmd = &cobra.Command{
	Use:     "list",
	Version: "0.0.1",
	RunE: versionListHandler,
}

func versionListHandler (cmd *cobra.Command, args []string) error {
	polvoAddress, err := cmd.Flags().GetString("polvo_address")
	if err != nil {
		return err
	}

	client, err := newPolvoClient(cmd.Context(), polvoAddress)
	if err != nil {
		return err
	}

	packageOrn, err := cmd.Flags().GetString("orn")
	if err != nil {
		return err
	}

	request := &v1.ListVersionsRequest{
		Orn: packageOrn,
	}

	stream, err := client.ListVersions(cmd.Context(), request)
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

		for _, ver := range response.GetVersions() {
			content, err := json.MarshalIndent(ver,  " ", " ")
			if err != nil {
				return err
			}

			fmt.Println(string(content))
		}
	}

}

func init() {
	versionListCmd.Flags().String("orn", "", "package orn")
}
