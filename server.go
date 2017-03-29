package main

import (
	"runtime"

	"github.com/CANARIA/canaria-api/config"
	"github.com/CANARIA/canaria-api/env"
	"github.com/CANARIA/canaria-api/router"
)

func main() {
	e := router.Init()
	runtime.GOMAXPROCS(runtime.NumCPU())
	env.SetUp()
	e.Logger.Fatal(e.Start(config.GetPort()))
}
