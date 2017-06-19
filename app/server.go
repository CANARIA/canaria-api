package main

import (
	"runtime"

	"github.com/CANARIA/canaria-api/env"
	"github.com/CANARIA/canaria-api/router"
)

func init() {
	app := router.Init()
	runtime.GOMAXPROCS(runtime.NumCPU())
	app.Logger.Fatal(app.Start(env.GetBind()))
}
