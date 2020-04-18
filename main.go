package main

import (
	"github.com/wolfplus2048/corona-api"
	"github.com/wolfplus2048/corona-demo/handler"
)

func main()  {
	app := corona.Default()
	app.Register(&handler.Handler{}, "handler")
	app.AddAcceptor(":3333")
	app.Configure(true, "example", nil)
	app.Start()
}