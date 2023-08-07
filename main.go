package main

import (
	"fmt"
	"github.com/ithaiq/thqgo/internal"
)

func main() {
	engine := internal.New()
	g := engine.Group("test")
	g.GET("/hello", func(c *internal.Context) {
		fmt.Fprintln(c.W, "hello thqgo")
	})
	engine.Run()
}
