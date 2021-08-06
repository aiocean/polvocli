package cmd

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/spf13/cobra"
	v1 "pkg.aiocean.dev/polvogo/aiocean/polvo/v1"
)

var packageListCmd = &cobra.Command{
	Use:     "list",
	Version: "0.0.1",
	RunE: packageListHandler,
}

func packageListHandler (cmd *cobra.Command, args []string) error {
	polvoAddress, err := cmd.Flags().GetString("polvo_address")
	if err != nil {
		return err
	}

	client, err := newPolvoClient(cmd.Context(), polvoAddress)
	if err != nil {
		return err
	}

	request := &v1.ListPackagesRequest{}

	stream, err := client.ListPackages(cmd.Context(), request)
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

		for _, pkg := range response.GetPackages() {
			content, err := json.MarshalIndent(pkg,  " ", " ")
			if err != nil {
				return err
			}

			fmt.Println(string(content))
		}
	}

}
