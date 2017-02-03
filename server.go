package main

import (
	"github.com/CANARIA/canaria-api/config"
	"github.com/CANARIA/canaria-api/router"
)

func main() {
	e := router.Init()
	e.Logger.Fatal(e.Start(config.GetPort()))
}
