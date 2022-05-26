package commander

import "context"

type CommandFn func(cmd *Command) error
type Commands []*Command

type Command struct {
	Name      string
	runFn     CommandFn
	prepareFn *CommandFn
	doneFn    *CommandFn
	errorFn   CommandFn
	error     error
	Ctx       context.Context
}

func (c *Command) Handle(ctx context.Context) error {
	c.Ctx = ctx

	if err := c.runPrepare(); err != nil {
		return c.runError(err)
	}

	if err := c.runFn(c); err != nil {
		return c.runError(err)
	}

	return c.runError(c.runDone())
}

func (c *Command) runPrepare() error {
	return c.fn(c.prepareFn)
}

func (c *Command) runDone() error {
	return c.fn(c.doneFn)
}
func (c *Command) runError(err error) error {
	if err == nil {
		return nil
	}

	c.error = err
	return c.fn(&c.errorFn)
}
func (c Command) Error() string {
	if c.error != nil {
		return c.error.Error()
	}
	return ``
}

func (c Command) GetError() error {
	return c.error
}

func (c *Command) fn(fn *CommandFn) error {
	if fn != nil {
		err := (*fn)(c)
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *Command) OnDone(fn CommandFn) *Command {
	c.doneFn = &fn
	return c
}

func (c *Command) OnRun(fn CommandFn) *Command {
	c.runFn = fn
	return c
}
func (c *Command) OnError(fn CommandFn) *Command {
	c.errorFn = fn
	return c
}

func (c *Command) OnPrepare(fn CommandFn) *Command {
	c.prepareFn = &fn
	return c
}

func NewCommand(name string) *Command {
	return &Command{
		Name:    name,
		errorFn: func(cmd *Command) error { return cmd.error },
	}
}
