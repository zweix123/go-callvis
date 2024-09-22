package main

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/zweix123/go-callvis/pkg/log"
)

type Flags struct {
	version bool
	debug   bool

	algo string
}

var flags Flags

func NewCommand() *cobra.Command {
	c := &cobra.Command{
		Use:   "go-callvis-modern",
		Short: "go-callvis-modern",
		Run: func(cmd *cobra.Command, args []string) {
			if flags.version {
				fmt.Println(Version())
				return
			}
			if flags.debug {
				log.SetValid()
				log.Log("debug is true")
			}

			log.Log("args: %v", args)

			// err := go_callvis_modern.GetInstance().Process()
			// if err != nil {
			// 	log.Log("error: %v", err)
			// }
		},
	}

	c.Flags().BoolVarP(&flags.version, "version", "v", false, "version")
	c.Flags().BoolVarP(&flags.debug, "debug", "d", false, "debug mode: print debug log")

	c.Flags().StringVarP(&flags.algo, "algo", "a", "static", "")

	return c
}

func main() {
	err := NewCommand().Execute()
	if err != nil {
		println(err.Error())
	}
}
