package main

import "fmt"

type command struct {
	name string
	args []string
}

type commands struct {
	handlers map[string]func(*state, command) error
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.handlers[name] = f
}

func (c *commands) run(s *state, cmd command) error {
	handlerName := cmd.name
	val, ok := c.handlers[handlerName]
	if !ok {
		return fmt.Errorf("no handlers found for '%s'", handlerName)
	}
	return val(s, cmd)
}
