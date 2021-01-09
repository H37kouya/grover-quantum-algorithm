package cmd

import (
	"github.com/spf13/cobra"
	"grover-quantum-search/pkg/usecase"
)

func New2qubitExecute() *cobra.Command {
	type Options struct{}

	var (
		o = &Options{}
	)

	cmd := &cobra.Command{
		Use:   "2qubit",
		Short: "A brief description of your command",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return validateParams(*o)
		},
		Run: func(cmd *cobra.Command, args []string) {
			usecase.Fixed2QubitUsecase()
		},
		SilenceErrors: true,
		SilenceUsage:  true,
	}

	return cmd
}
