package app

import (
	"runtime"

	"github.com/CANARIA/canaria-api/core/router"
	"net/http"
)

func init() {
	app := router.Init()
	runtime.GOMAXPROCS(runtime.NumCPU())
	http.Handle("/", app)
}
