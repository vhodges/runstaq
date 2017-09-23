package main

import (
	"fmt"

	"github.com/abiosoft/ishell"
)

func listCmd() *ishell.Cmd {
	return &ishell.Cmd{
		Name:      "list",
		Help:      "List staq components and their statuses",
		Completer: Completer,
		Func: func(c *ishell.Context) {

			pattern := "*/*" // Default to all

			if len(c.Args) > 0 {
				pattern = c.Args[0]
			}

			for _, procfile := range AppStaq.Procfiles {
				for _, proc := range procfile.Procs {
					if Glob(pattern, fmt.Sprintf("%s/%s", procfile.Name, proc.Name)) {
						c.Printf("%s %s : %10s %s\n",
							procfile.Name,
							procfile.Status(),
							proc.Name,
							proc.Status())
					}
				}

				c.Printf("\n")
				c.Printf("\n")
			}
		},
	}
}
