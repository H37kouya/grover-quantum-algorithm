package cmd

import (
	"github.com/spf13/cobra"
	"grover-quantum-search/pkg/usecase"
)

func NewNqubitCsvExecute() *cobra.Command {
	type Options struct{
		Optint int    `validate:"min=1,max=30"`
	}

	var (
		o = &Options{}
	)

	cmd := &cobra.Command{
		Use:   "nqubit-csv",
		Short: "A brief description of your command",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return validateParams(*o)
		},
		Run: func(cmd *cobra.Command, args []string) {
			usecase.FixedNQubitCsvUsecase(o.Optint)
		},
		SilenceErrors: true,
		SilenceUsage:  true,
	}

	cmd.Flags().IntVarP(&o.Optint, "nqubit", "n", 1, "nqubit")

	return cmd
}

func NewNqubitTimesExecute() *cobra.Command {
	type Options struct{
		Optint int    `validate:"min=2,max=30"`
	}

	var (
		o = &Options{}
	)

	cmd := &cobra.Command{
		Use:   "nqubit-times",
		Short: "A brief description of your command",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return validateParams(*o)
		},
		Run: func(cmd *cobra.Command, args []string) {
			usecase.FixedNQubitTimesUsecase(o.Optint)
		},
		SilenceErrors: true,
		SilenceUsage:  true,
	}

	cmd.Flags().IntVarP(&o.Optint, "nqubit", "n", 2, "nqubit")

	return cmd
}

func NewNqubitTimesAllExecute() *cobra.Command {
	type Options struct{}

	var (
		o = &Options{}
	)

	cmd := &cobra.Command{
		Use:   "nqubit-times-all",
		Short: "A brief description of your command",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return validateParams(*o)
		},
		Run: func(cmd *cobra.Command, args []string) {
			usecase.FixedNQubitTimesAllUsecase()
		},
		SilenceErrors: true,
		SilenceUsage:  true,
	}

	return cmd
}
