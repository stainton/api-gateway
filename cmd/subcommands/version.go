package subcommands

import (
	"fmt"

	"github.com/spf13/cobra"
)

const _VERSION string = "v1.0.0"

func NewSubcmdVersion() *cobra.Command {
	cmdVersion := &cobra.Command{
		Use:     "version",
		Short:   "version of gateway command",
		Long:    "version of gateway command",
		Example: "gateway version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("%s %s\n", cmd.Parent().Name(), _VERSION)
		},
	}
	return cmdVersion
}
