package subcommands

import (
	"github.com/spf13/cobra"
	"github.com/stainton/api-gateway/internal"
)

func NewSubcmdServe() *cobra.Command {
	var connectionString string
	cmdServe := &cobra.Command{
		Use:     "serve",
		Short:   "run gateway server",
		Long:    "run gateway server",
		Example: "gateway serve",
		Run: func(cmd *cobra.Command, args []string) {
			internal.Run(connectionString)
		},
	}
	cmdServe.Flags().StringVarP(&connectionString, "addr", "a", ":8080", "address gateway server listening")
	return cmdServe
}
