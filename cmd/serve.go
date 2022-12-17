package cmd

import (
	"github.com/spf13/cobra"

	"ecommerce/customer/internal/app"
	"ecommerce/customer/pkg/logger"
)

func serveCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "serve",
		Short: "Start a new grpc server",
		Run: func(cmd *cobra.Command, args []string) {
			logger.Initialize()
			app.RunGRPCServer()
		},
	}
}
