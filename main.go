package main

import (
	"dig/app"
	"dig/clients"
	"dig/internal/config"
	"dig/internal/helper"
	"dig/internal/logger"

	"go.uber.org/dig"
)

func main() {
	c := dig.New()

	// Provides all dependencies
	err := helper.Provide(c,
		config.New,
		logger.New,
		clients.Module,
		app.New,
	)
	if err != nil {
		panic(err)
	}

	// Call your program entrypoint
	err = c.Invoke(func(a *app.App) {
		a.Run()
	})
	if err != nil {
		panic(err)
	}
}
