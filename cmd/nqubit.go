package cmd

import (
	"github.com/spf13/cobra"
	"grover-quantum-search/pkg/domain/valueObject"
	"grover-quantum-search/pkg/usecase"
)

func NewNqubitCsvExecute() *cobra.Command {
	type Options struct {
		Optint int `validate:"min=1,max=30"`
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
	type Options struct {
		Optint int `validate:"min=2,max=30"`
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
	type Options struct {
		Min int `validate:"min=2,max=30"`
		Max int `validate:"min=2,max=30"`
	}

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
			usecase.FixedNQubitTimesAllUsecase(o.Min, o.Max)
		},
		SilenceErrors: true,
		SilenceUsage:  true,
	}

	cmd.Flags().IntVarP(&o.Min, "min", "s", 4, "最小値")
	cmd.Flags().IntVarP(&o.Max, "max", "e", 25, "最大値")

	return cmd
}

func NewRandomNqubitCsvExecute() *cobra.Command {
	type Options struct {
		Optint   int `validate:"min=1,max=30"`
		Loop     int `validate:"min=1,max=100000"`
		PlusReal float64
		PlusImag float64
	}

	var (
		o = &Options{}
	)

	cmd := &cobra.Command{
		Use:   "nqubit-random-csv",
		Short: "A brief description of your command",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return validateParams(*o)
		},
		Run: func(cmd *cobra.Command, args []string) {
			usecase.RandomNQubitCsvUsecase(
				o.Optint,
				valueObject.Qubit(complex(o.PlusReal, o.PlusImag)),
				o.Loop,
			)
		},
		SilenceErrors: true,
		SilenceUsage:  true,
	}

	cmd.Flags().IntVarP(&o.Optint, "nqubit", "n", 1, "nqubit")
	cmd.Flags().IntVarP(&o.Loop, "loop", "l", 512, "nqubit")
	cmd.Flags().Float64VarP(&o.PlusReal, "plus-real", "r", 0.0, "nqubit")
	cmd.Flags().Float64VarP(&o.PlusImag, "plus-imag", "i", 0.0, "nqubit")

	return cmd
}

func NewRandomNqubitTimesExecute() *cobra.Command {
	type Options struct {
		Optint   int `validate:"min=2,max=30"`
		OptCount int `validate:"min=1"`
		PlusReal float64
		PlusImag float64
	}

	var (
		o = &Options{}
	)

	cmd := &cobra.Command{
		Use:   "nqubit-random-times",
		Short: "A brief description of your command",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return validateParams(*o)
		},
		Run: func(cmd *cobra.Command, args []string) {
			usecase.RandomNQubitTimesCountUsecase(
				o.Optint,
				o.OptCount,
				o.PlusReal,
				o.PlusImag,
			)
		},
		SilenceErrors: true,
		SilenceUsage:  true,
	}

	cmd.Flags().IntVarP(&o.Optint, "nqubit", "n", 2, "nqubit")
	cmd.Flags().IntVarP(&o.OptCount, "count", "c", 1, "nqubit")
	cmd.Flags().Float64VarP(&o.PlusReal, "plus-real", "r", 0.0, "nqubit")
	cmd.Flags().Float64VarP(&o.PlusImag, "plus-imag", "i", 0.0, "nqubit")

	return cmd
}
