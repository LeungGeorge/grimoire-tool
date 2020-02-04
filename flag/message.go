package flag

import (
	"github.com/urfave/cli"
)

// Message git commit message
var Message string

// FlagMessage ...
var FlagMessage = cli.StringFlag{
	Name:        "message, m",                // flag 及其简写
	Value:       "auto commit",               // 默认值
	Usage:       "the commit message of git", // 描述信息
	Destination: &Message,                    // flag 接受者
}
