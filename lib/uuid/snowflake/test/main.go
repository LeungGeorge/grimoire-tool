package main

import (
	"fmt"

	"github.com/LeungGeorge/grimoire-tool/lib/uuid/snowflake"
)

func main() {
	g, err := snowflake.NewGUID(1)
	if err != nil {
		panic(err)
	}

	for i := 0; i < 1000; i++ {
		fmt.Println(g.NextID())
	}
}
