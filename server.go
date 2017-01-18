package main

import "github.com/CANARIA/canaria-api/router"
import "github.com/CANARIA/canaria-api/config"

func main() {
	e := router.Init()
	e.Logger.Fatal(e.Start(config.GetPort()))
}
