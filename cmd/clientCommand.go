package cmd

import (
	"go-deamon-sample/client"

	"github.com/spf13/cobra"
)

type Opt struct {
	standAlone bool
}

func NewClientCommand() *cobra.Command {
	o := &Opt{}

	cmd := cobra.Command{
		Use: "client",
		Run: func(cmd *cobra.Command, args []string) {
			client.Do(o.standAlone)
		},
	}

	cmd.Flags().BoolVarP(&o.standAlone, "standalone", "", false, "stand-alone option")

	return &cmd
}
