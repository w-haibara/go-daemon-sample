package cmd

import (
	"go-deamon-sample/deamon"

	"github.com/spf13/cobra"
)

var (
	deamonCmd = cobra.Command{
		Use: "deamon",
		Run: func(cmd *cobra.Command, args []string) {
			deamon.Do()
		},
	}
)
