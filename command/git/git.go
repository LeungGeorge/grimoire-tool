package git

import (
	"github.com/urfave/cli"
)

const (
	statusEmpty   = 0x00
	statusAHead   = 0x01
	statusHasFile = 0x02
)

// Git 操作封装
var Git = cli.Command{
	Name:     "git",
	Usage:    "git commands: sync.",
	Category: "git",
	Subcommands: []cli.Command{
		{
			Name:   "sync",
			Usage:  "sync current git repository.",
			Action: sync,
		},
	},
}
