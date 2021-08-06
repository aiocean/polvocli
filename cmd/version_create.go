package cmd

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/spf13/cobra"
	v1 "pkg.aiocean.dev/polvogo/aiocean/polvo/v1"
)

var versionCreateCmd = &cobra.Command{
	Use:     "create",
	Version: "0.0.1",
	RunE: versionVersionHandler,
}

func versionVersionHandler (cmd *cobra.Command, args []string) error {
	polvoAddress, err := cmd.Flags().GetString("polvo_address")
	if err != nil {
		return err
	}

	versionName, err := cmd.Flags().GetString("name")
	if err != nil {
		return err
	}

	manifestUrl, err := cmd.Flags().GetString("manifest_url")
	if err != nil {
		return err
	}

	packageOrn, err := cmd.Flags().GetString("package_orn")
	if err != nil {
		return err
	}

	weight, err := cmd.Flags().GetUint32("weight")
	if err != nil {
		return err
	}

	client, err := newPolvoClient(cmd.Context(), polvoAddress)
	if err != nil {
		return err
	}

	request := &v1.CreateVersionRequest{
		PackageOrn: packageOrn,
		Version: &v1.Version{
			Name:        versionName,
			ManifestUrl: manifestUrl,
			Weight: weight,
		},
	}

	stream, err := client.CreateVersion(cmd.Context(), request)
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

		content, err := json.MarshalIndent(response.GetVersion(), " ", " ")
		if err != nil {
			return err
		}

		fmt.Println(string(content))
	}

}

func init() {
	versionCreateCmd.Flags().String("name", "", "version name")
	versionCreateCmd.Flags().String("manifest_url", "", "version's manifest url")
	versionCreateCmd.Flags().String("package_orn", "", "package orn")
	versionCreateCmd.Flags().Uint32("weight", 0, "delivery weight")
}
