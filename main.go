package main

import (
	"corona-demo/demo"
	"github.com/wolfplus2048/corona-api"
	"log"
	"plugin"
	"reflect"
)
var (
	app corona.App
)
func init()  {
	lib, err := plugin.Open("./corona.so")
	if err != nil {
		log.Printf("plugin open err :%v", err)
		panic(err)
	}
	sym, err := lib.Lookup("app")
	if err != nil {
		log.Printf("lib Lookup err :%v", err)
		panic(err)
	}
	a, ok := sym.(corona.App)
	if !ok {
		log.Printf("expecting corona app but %s", reflect.TypeOf(a).Name())
		panic("expecting corona app")
	}
	app = a
}
func main()  {
	app.Register(&demo.Handler{}, "demo")
	app.AddAcceptor(":3333")
	app.Configure(true, "demo", nil)
	app.Start()
}