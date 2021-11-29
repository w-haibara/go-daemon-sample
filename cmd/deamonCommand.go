package cmd

import (
	"go-deamon-sample/deamon"

	"github.com/spf13/cobra"
)

func NewDeamonCommand() *cobra.Command {
	cmd := cobra.Command{
		Use: "deamon",
		Run: func(cmd *cobra.Command, args []string) {
			deamon.Do()
		},
	}
	return &cmd
}
