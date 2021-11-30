package cmd

import (
	"fmt"
	"go-daemon-sample/client"
	"log"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

type Opt struct {
	standAlone bool
}

var (
	o = &Opt{}
)

func init() {
	clientCmd.AddCommand(&clientAddCmd)
	clientCmd.AddCommand(&clientSubCmd)
	clientCmd.PersistentFlags().BoolVarP(&o.standAlone, "standalone", "", false, "stand-alone option")
}

var (
	clientCmd = cobra.Command{
		Use: "client",
		Run: func(cmd *cobra.Command, args []string) {
			if err := cmd.Help(); err != nil {
				log.Fatalln(err)
			}
			os.Exit(0)
		},
	}

	clientAddCmd = cobra.Command{
		Use: "add",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 2 {
				return fmt.Errorf(`"%s" requires 2 arguments`, cmd.Name())
			}
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			a, err := strconv.Atoi(args[0])
			if err != nil {
				log.Fatalln(err)
			}

			b, err := strconv.Atoi(args[1])
			if err != nil {
				log.Fatalln(err)
			}
			client.Add(a, b, o.standAlone)
		},
	}

	clientSubCmd = cobra.Command{
		Use: "sub",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 2 {
				return fmt.Errorf(`"%s" requires 2 arguments`, cmd.Name())
			}
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			a, err := strconv.Atoi(args[0])
			if err != nil {
				log.Fatalln(err)
			}

			b, err := strconv.Atoi(args[1])
			if err != nil {
				log.Fatalln(err)
			}
			client.Sub(a, b, o.standAlone)
		},
	}
)
