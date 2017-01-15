package main

import "github.com/CANARIA/canaria-api/router"
import "github.com/canaria/canaria-api/config"

func main() {
	e := router.Init()
	e.Logger.Fatal(e.Start(config.GetPort()))
}
