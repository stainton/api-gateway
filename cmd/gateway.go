package cmd

import (
	"github.com/spf13/cobra"
	"github.com/stainton/api-gateway/cmd/subcommands"
)

func NewCmdGateway() *cobra.Command {
	gateCmd := &cobra.Command{
		Use:     "gateway [command]",
		Short:   "gateway of apiserver",
		Long:    "gateway of apiserver",
		Example: "gateway version",
		// Run:     func(cmd *cobra.Command, args []string) {},
	}
	// gateCmd.Flags().StringVarP()
	gateCmd.AddCommand(subcommands.NewSubcmdVersion())
	gateCmd.AddCommand(subcommands.NewSubcmdServe())
	return gateCmd
}
