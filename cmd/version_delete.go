package cmd

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/spf13/cobra"
	v1 "pkg.aiocean.dev/polvogo/aiocean/polvo/v1"
)

var versionDeleteCmd = &cobra.Command{
	Use:     "delete",
	Version: "0.0.1",
	RunE: versionDeleteHandler,
}

func versionDeleteHandler (cmd *cobra.Command, args []string) error {
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

	request := &v1.DeleteVersionRequest{
		Orn: orn,
	}

	stream, err := client.DeleteVersion(cmd.Context(), request)
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

		content, err := json.MarshalIndent(response.GetMessage(), " ", " ")
		if err != nil {
			return err
		}

		fmt.Println(string(content))
	}

}

func init() {
	versionDeleteCmd.Flags().String("orn", "", "version orn")
}
