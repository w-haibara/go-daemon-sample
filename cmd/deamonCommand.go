package cmd

import (
	"go-daemon-sample/daemon"

	"github.com/spf13/cobra"
)

var (
	daemonCmd = cobra.Command{
		Use: "daemon",
		Run: func(cmd *cobra.Command, args []string) {
			daemon.Do()
		},
	}
)
