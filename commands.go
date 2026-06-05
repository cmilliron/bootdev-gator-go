package main

import "fmt"

type command struct {
	Name	string
	Args	[]string
}

type commands struct {
	registry 	map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	f, exist := c.registry[cmd.Name]
	if exist == false {
		return fmt.Errorf("Command does not exist.\n")
	}
	
	return f(s, cmd)
}


func (c *commands) register(name string, f func(s *state, cmd command) error) error {
	c.registry[name] = f
	return nil
}