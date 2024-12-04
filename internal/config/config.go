package config

import "go.uber.org/dig"

type Config struct {
	Prefix      string
	ClientName  string
	StorageName string
}

// example of dig.Out
type NewOutput struct {
	dig.Out
	Cfg *Config
}

func New() NewOutput {
	return NewOutput{
		Cfg: &Config{
			Prefix:      "[foo] ",
			ClientName:  "MyClient",
			StorageName: "MyStorage",
		},
	}
}
