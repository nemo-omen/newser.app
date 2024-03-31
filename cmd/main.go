package main

import (
	"context"

	"newser.app/internal/app"
)

func main() {
	app.NewApp(context.Background()).Start()
}
