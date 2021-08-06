package cmd

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/spf13/cobra"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	aiocean_polvo_v1 "pkg.aiocean.dev/polvogo/aiocean/polvo/v1"
	v1 "pkg.aiocean.dev/polvoservice/pkg/client/v1"
)

var versionUpdateCmd = &cobra.Command{
	Use:     "update",
	Short:   "Update version build.",
	Version: "0.0.1",
	RunE:    versionUpdateRun,
}

func init() {
	versionUpdateCmd.Flags().String("orn", "", "version orn")
	versionUpdateCmd.Flags().String("name", "", "version name")
	versionUpdateCmd.Flags().Uint32("weight", 0, "version's weight")
	versionUpdateCmd.Flags().String("manifest_url", "", "Entry point url")

	versionUpdateCmd.MarkFlagRequired("orn")
}

func versionUpdateRun(cmd *cobra.Command, args []string) error {
	polvoAddress, err := cmd.Parent().PersistentFlags().GetString("polvo_address")
	if err != nil {
		return err
	}

	client, err := v1.NewClient(cmd.Context(),  &v1.Config{
		Address: polvoAddress,
	})
	if err != nil {
		return err
	}

	request := &aiocean_polvo_v1.UpdateVersionRequest{
		Version:    &aiocean_polvo_v1.Version{},
		FieldMask: &fieldmaskpb.FieldMask{},
	}

	if orn, err := cmd.Flags().GetString("orn"); err != nil {
		return err
	} else if orn != "" {
		request.Orn = orn
	}

	if manifestUrl, err := cmd.Flags().GetString("manifest_url"); err != nil {
		return err
	} else if manifestUrl != "" {
		request.Version.ManifestUrl = manifestUrl
		if err := request.FieldMask.Append(request, "version.manifest_url"); err != nil {
			return err
		}
	}

	if weight, err := cmd.Flags().GetUint32("weight"); err != nil {
		return err
	} else if weight > 0{
		request.Version.Weight = weight
		if err := request.FieldMask.Append(request, "version.weight"); err != nil {
			return err
		}
	}

	stream, err := client.UpdateVersion(cmd.Context(), request)
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
