package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/zweix123/go-callvis/internal/go_callvis_modern"
	"github.com/zweix123/go-callvis/pkg/log"
)

type Flags struct {
	version bool
	debug   bool

	algo string

	test bool
}

var flags Flags

func NewCommand() *cobra.Command {
	c := &cobra.Command{
		Use:   "go-callvis-modern",
		Short: "go-callvis-modern",
		Run: func(cmd *cobra.Command, args []string) {
			// develop config
			// flags.debug = true
			// flags.algo = "static"

			if flags.version {
				fmt.Println(Version())
				return
			}
			if flags.debug {
				log.SetValid()
				log.Log("debug is true")
			}

			log.Log("args: %#v", args)

			if len(args) == 0 {
				// 没有任何参数, 默认分析执行命令的路径
				go_callvis_modern.GetInstance().Dir = ""
			} else if len(args) == 1 {
				// 只有一个参数, go-callvis-modern最多一个参数, 所以这个参数就是分析的路径
				go_callvis_modern.GetInstance().Dir = args[0]
			} else {
				// 参数太多, 直接退出
				log.Log("error: too many arguments, only one argument is allowed, which is the path to analyze")
				return
			}

			go_callvis_modern.GetInstance().Algo = go_callvis_modern.Algorithm(flags.algo)

			go_callvis_modern.GetInstance().IncludeTest = flags.test

			log.Log("argument: %s", go_callvis_modern.GetInstance().Argument())

			err := go_callvis_modern.GetInstance().Process()
			if err != nil {
				log.Log("error: %v", err)

				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
		},
	}

	c.Flags().BoolVarP(&flags.version, "version", "v", false, "version")
	c.Flags().BoolVarP(&flags.debug, "debug", "d", false, "debug mode: print debug log")

	c.Flags().StringVarP(&flags.algo, "algo", "a", "static", "")

	c.Flags().BoolVarP(&flags.test, "test", "t", false, "analysis include test file")

	return c
}

func main() {
	err := NewCommand().Execute()
	if err != nil {
		println(err.Error())
	}
}
