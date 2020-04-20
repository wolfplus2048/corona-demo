package main

import (
	"github.com/wolfplus2048/corona-api"
	"github.com/wolfplus2048/corona-demo/handler"
	"net/http"
)

func main()  {
	http.Handle("/web/", http.StripPrefix("/web/", http.FileServer(http.Dir("web"))))
	go http.ListenAndServe(":3334", nil)

	app := corona.Default()
	app.Register(&handler.Handler{}, "handler")
	app.AddAcceptor(":3333")
	app.Configure(true, "demo", nil)
	app.Start()
}