package main

import (
	"runtime"

	"github.com/CANARIA/canaria-api/core/env"
	"github.com/CANARIA/canaria-api/core/router"
)

func main() {
	app := router.Init()
	runtime.GOMAXPROCS(runtime.NumCPU())
	app.Logger.Fatal(app.Start(env.GetBind()))
}
