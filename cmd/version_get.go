package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	v1 "pkg.aiocean.dev/polvogo/aiocean/polvo/v1"
)

var versionGetCmd = &cobra.Command{
	Use:     "get",
	Version: "0.0.1",
	RunE: versionGetHandler,
}

func versionGetHandler (cmd *cobra.Command, args []string) error {
	polvoAddress, err := cmd.Flags().GetString("polvo_address")
	if err != nil {
		return err
	}

	versionOrn, err := cmd.Flags().GetString("orn")
	if err != nil {
		return err
	}

	client, err := newPolvoClient(cmd.Context(), polvoAddress)
	if err != nil {
		return err
	}

	request := &v1.GetVersionRequest{
		Orn: versionOrn,
	}

	response, err := client.GetVersion(cmd.Context(), request)
	if err != nil {
		return err
	}

	content, err := json.MarshalIndent(response.GetVersion(), " ", " ")
	if err != nil {
		return err
	}

	fmt.Println(string(content))

	return nil
}

func init() {
	versionGetCmd.Flags().String("orn", "", "version orn")
}
