package app

import (
	"dig/clients/cloudprovider"
	"dig/clients/storage"
	"dig/internal/logger"
)

type App struct {
	logger        logger.Interface
	cloudprovider *cloudprovider.Client
	storage       *storage.Client
}

func (a *App) Run() error {
	a.logger.Print("Running app with cloud provider: " + a.cloudprovider.Name() +
		" and storage provider: " + a.storage.Name())
	return nil
}

func New(l logger.Interface, c *cloudprovider.Client, s *storage.Client) *App {
	return &App{
		logger:        l,
		cloudprovider: c,
		storage:       s,
	}
}
