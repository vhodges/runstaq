package main

import (
	"github.com/abiosoft/ishell"
)

func startCmd() *ishell.Cmd {
	return &ishell.Cmd{
		Name:      "run",
		Help:      "run [component] executing the staq components or a specific component",
		Completer: Completer,
		Func: func(c *ishell.Context) {
			if len(c.Args) == 0 {
				for _, procfile := range AppStaq.Procfiles {
					procfile.Start()
				}
			} else {
				procfile := AppStaq.module(c.Args[0])
				if procfile != nil {
					procfile.Start()
				} else {
					c.Printf("No such module: %s\n", c.Args[0])
				}
			}
		},
	}
}
