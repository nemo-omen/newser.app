package main

import (
	"context"
	"os"

	"newser.app/config"
	"newser.app/internal/app"
)

func main() {
	env := os.Getenv("env")
	if env == "" {
		env = "local"
	}
	cfg, err := config.LoadConfig(env)
	if err != nil {
		panic(err)
	}
	app.NewApp(context.Background(), cfg).Start()
}
