package command

import (
	"fmt"
)


type Command[T any] interface {
	Usage() (string)
	Execute([]string, *T)
}


type Group[T any] struct {
	commands map[string]Command[T]
}

func (c *Group[T]) Register(name string, cmd Command[T]) {
	if c.commands == nil {
		c.commands = make(map[string]Command[T])
	}
	c.commands[name] = cmd
}

func (g *Group[T]) Usage() string {
	return "group"
}

func (g *Group[T]) Execute(f []string, arg *T) {
	for name, cmd := range g.commands {
		if name == f[0] {
			cmd.Execute(f[1:], arg)
			return
		}
	}

	fmt.Println(g.Usage())
}
