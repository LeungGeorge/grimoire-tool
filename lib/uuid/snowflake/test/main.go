package main

import "fmt"
import "github.com/LeungGeorge/grit/lib/uuid/snowflake"

func main() {
	g, err := snowflake.NewGUID(1)
	if err != nil {
		panic(err)
	}

	for i := 0; i < 1000; i++ {
		fmt.Println(g.NextID())
	}
}
