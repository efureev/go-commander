package commander

import (
	"context"
)

type Commander struct {
	commands Commands
	Ctx      context.Context
	error    error
}

func (c *Commander) Add(cmd ...*Command) *Commander {
	c.commands = append(c.commands, cmd...)

	return c
}

func (c *Commander) Run() *Commander {
	for _, cmd := range c.commands {

		if err := cmd.Handle(c.Ctx); err != nil {
			c.error = err
			return c
		}

		c.Ctx = cmd.Ctx
	}

	return c
}

func (c Commander) HasError() bool {
	return c.error != nil
}

func (c Commander) Error() error {
	return c.error
}

func (c *Commander) updateCtx(ctx context.Context) {
	c.Ctx = ctx
}

func NewCommander() *Commander {
	return &Commander{
		Ctx: context.Background(),
	}
}
