package cmd

import (
	"go-deamon-sample/client"

	"github.com/spf13/cobra"
)

func NewClientCommand() *cobra.Command {
	cmd := cobra.Command{
		Use: "client",
		Run: func(cmd *cobra.Command, args []string) {
			client.Do()
		},
	}
	return &cmd
}
