package helper

import (
	"go.uber.org/dig"
)

type ModuleOutput func() []any

func Module(p ...any) ModuleOutput {
	return func() []any {
		return p
	}
}

func Provide(c *dig.Container, p ...any) error {
	for _, constructor := range p {
		switch v := constructor.(type) {
		case ModuleOutput:
			// Handle Annotation type
			if err := Provide(c, v()...); err != nil {
				return err
			}
		default:
			if err := c.Provide(constructor); err != nil {
				return err
			}
		}
	}
	return nil
}
