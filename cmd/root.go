package cmd

import (
	"github.com/spf13/cobra"
)

func Execute() {
	var cmd = &cobra.Command{
		Use: "go-deamon-sample",
	}

	cmd.AddCommand(NewDeamonCommand())
	cmd.AddCommand(NewClientCommand())
	cobra.CheckErr(cmd.Execute())
}
