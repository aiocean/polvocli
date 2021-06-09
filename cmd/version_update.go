package cmd

import (
	"errors"

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
	versionUpdateCmd.Flags().String("orn", "", "orn")
	versionUpdateCmd.Flags().String("display_name", "", "version's display name")
	versionUpdateCmd.Flags().String("git_ref", "", "git ref")
	versionUpdateCmd.Flags().String("build_name", "", "build name")
	versionUpdateCmd.Flags().String("entry_point_url", "", "Entry point url")

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
		UpdateMask: &fieldmaskpb.FieldMask{},
	}

	if orn, err := cmd.Flags().GetString("orn"); err != nil {
		return err
	} else if orn == "" {
		return errors.New("orn is required")
	} else {
		request.Version.Orn = orn
	}

	if displayName, err := cmd.Flags().GetString("display_name"); err != nil {
		return err
	} else if displayName != "" {
		request.Version.DisplayName = displayName
		request.UpdateMask.Paths = append(request.UpdateMask.Paths, "version.display_name")
	}

	if gitRef, err := cmd.Flags().GetString("git_ref"); err != nil {
		return err
	} else if gitRef != "" {
		request.Version.GitRef = gitRef
		request.UpdateMask.Paths = append(request.UpdateMask.Paths, "version.git_ref")
	}

	if entryPointUrl, err := cmd.Flags().GetString("entry_point_url"); err != nil {
		return err
	} else if entryPointUrl != "" {
		request.Version.EntryPointUrl = entryPointUrl
		request.UpdateMask.Paths = append(request.UpdateMask.Paths, "version.entry_point_url")
	}

	if buildName, err := cmd.Flags().GetString("build_name"); err != nil {
		return err
	} else if buildName != "" {
		request.Version.BuildName = buildName
		request.UpdateMask.Paths = append(request.UpdateMask.Paths, "version.build_name")
	}

	response, err := client.UpdateVersion(cmd.Context(), request)
	if err != nil {
		return err
	}

	cmd.Println("Done: " + response.GetVersion().GetOrn())

	return nil
}
