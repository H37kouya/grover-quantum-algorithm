package cmd

import (
	"fmt"
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Cmd struct{}

var cfgFile string

// NewCmd
func NewCmd() Cmd {
	return Cmd{}
}

func (c Cmd) Execute() {
	cmd := c.NewCmdRoot()
	cmd.SetOut(os.Stdout)
	if err := cmd.Execute(); err != nil {
		cmd.SetOut(os.Stderr)
		cmd.Println(err)
		os.Exit(1)
	}
}

func (c Cmd) NewCmdRoot() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "grover",
		Short: "A brief description of your application",
	}
	cobra.OnInitialize(initConfig)

	cmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cmd-test.yaml)")
	cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	cmd.AddCommand(New2qubitExecute())
	cmd.AddCommand(NewNqubitCsvExecute())
	cmd.AddCommand(NewNqubitTimesExecute())
	cmd.AddCommand(NewNqubitTimesAllExecute())
	cmd.AddCommand(NewRandomNqubitTimesExecute())
	cmd.AddCommand(NewRandomNqubitCsvExecute())
	return cmd
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		viper.AddConfigPath(home)
		viper.SetConfigName(".cmd-test")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
