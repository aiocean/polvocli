package cmd

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/spf13/cobra"
	v1 "pkg.aiocean.dev/polvogo/aiocean/polvo/v1"
)

var packageCreateCmd = &cobra.Command{
	Use:     "create",
	Version: "0.0.1",
	RunE: packagePackageHandler,
}

func packagePackageHandler (cmd *cobra.Command, args []string) error {
	polvoAddress, err := cmd.Flags().GetString("polvo_address")
	if err != nil {
		return err
	}

	packageName, err := cmd.Flags().GetString("name")
	if err != nil {
		return err
	}

	maintainer, err := cmd.Flags().GetString("maintainer")
	if err != nil {
		return err
	}

	client, err := newPolvoClient(cmd.Context(), polvoAddress)
	if err != nil {
		return err
	}

	request := &v1.CreatePackageRequest{
		Package: &v1.Package{
			Name: packageName,
			Maintainer: maintainer,
		},
	}

	stream, err := client.CreatePackage(cmd.Context(), request)
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

		content, err := json.MarshalIndent(response.GetPackage(), " ", " ")
		if err != nil {
			return err
		}

		fmt.Println(string(content))
	}

}

func init() {
	packageCreateCmd.Flags().String("name", "", "package name")
	packageCreateCmd.Flags().String("maintainer", "", "package's Maintainer")
}
