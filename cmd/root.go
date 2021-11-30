package cmd

import (
	"github.com/spf13/cobra"
)

func Execute() {
	var cmd = &cobra.Command{
		Use: "go-daemon-sample",
	}

	cmd.AddCommand(&daemonCmd)
	cmd.AddCommand(&clientCmd)
	cobra.CheckErr(cmd.Execute())
}
