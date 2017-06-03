package main

import (
	"github.com/abiosoft/ishell"
)

func stopCmd() *ishell.Cmd {
	return &ishell.Cmd{
		Name:      "shutdown",
		Help:      "shutdown [component] stops executing of the staq components or a specific component",
		Completer: Completer,
		Func: func(c *ishell.Context) {
			if len(c.Args) == 0 {
				for _, procfile := range AppStaq.Procfiles {
					procfile.Stop()
				}
			} else {
				procfile := AppStaq.module(c.Args[0])
				if procfile != nil {
					procfile.Stop()
				} else {
					c.Printf("No such module: %s\n", c.Args[0])
				}
			}
		},
	}
}
