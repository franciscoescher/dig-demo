package logger

import (
	"dig/internal/config"
	"log"
	"os"

	"go.uber.org/dig"
)

// example of providing interface
type Interface interface {
	Print(v ...any)
}

// example of dig.In
type NewParams struct {
	dig.In
	Cfg *config.Config
}

func New(p NewParams) Interface {
	return log.New(os.Stdout, p.Cfg.Prefix, 0)
}
