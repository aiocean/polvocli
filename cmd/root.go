package cmd

import (
	"context"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	v1 "pkg.aiocean.dev/polvogo/aiocean/polvo/v1"
	"pkg.aiocean.dev/serviceutil/grpcutils"
)

var RootCmd = &cobra.Command{
	Use:     "polvocli",
	Short:   "Một công cụ hỗt rợ làm việc với các polvo.",
	Version: "0.0.1",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Usage()
	},
}

func init() {
	RootCmd.AddCommand(versionCmd)
	RootCmd.AddCommand(packageCmd)
}

func newPolvoClient(ctx context.Context, address string) (v1.PolvoServiceClient, error) {
	conn, err := grpcutils.NewGrpcConnection(ctx, address)
	if err != nil {
		return nil, errors.WithMessage(err, "failed to dial")
	}

	client := v1.NewPolvoServiceClient(conn)

	return client, nil
}
