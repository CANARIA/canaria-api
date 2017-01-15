package main

import (
	"github.com/CANARIA/canaria-api/config"
	"github.com/CANARIA/canaria-api/router"

	"github.com/labstack/echo/engine/standard"
)

func main() {
	router := router.Init()
	router.Run(standard.New(config.GetPort()))
}
