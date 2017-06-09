package main

import (
	"github.com/abiosoft/ishell"
)

func listCmd() *ishell.Cmd {
	return &ishell.Cmd{
		Name:      "list",
		Help:      "List staq components and their statuses",
		Completer: Completer,
		Func: func(c *ishell.Context) {

			if len(c.Args) == 0 {
				for _, procfile := range AppStaq.Procfiles {
					c.Printf("%-15s %s\n", procfile.Name, procfile.Status())

					for _, proc := range procfile.Procs {
						c.Printf("%10s %s\n", proc.Name, proc.Status())
					}

					c.Printf("\n")
					c.Printf("\n")
				}
			} else {
				procfile := AppStaq.module(c.Args[0])
				if procfile != nil {
					c.Printf("%-15s %s\n", procfile.Name, procfile.Status())
					for _, proc := range procfile.Procs {
						c.Printf("%10s %s\n", proc.Name, proc.Status())
					}
				} else {
					c.Printf("No such module: %s\n", c.Args[0])
				}
			}
		},
	}
}
