package main

import (
	"flag"
	"net/http"

	"newser.app/cmd/web/core"
)

func main() {
	addr := flag.String("addr", "4321", "HTTP network address")
	dev := flag.Bool("dev", false, "Whether to run the server in development mode")
	flag.Parse()

	app := &core.App{
		Logger: core.Logger(*dev),
	}

	app.Logger.Info("starting server", "addr", *addr)
	app.Logger.Debug("dev mode", "dev", *dev)
	err := http.ListenAndServe(":"+*addr, app.Routes())
	app.Logger.Error(err.Error())
}
