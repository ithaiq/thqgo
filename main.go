package main

import (
	"fmt"
	"github.com/ithaiq/thqgo/internal"
	"net/http"
)

func main() {
	engine := internal.New()
	engine.Add("/hello", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer, "hello thqgo")
	})
	engine.Run()
}
