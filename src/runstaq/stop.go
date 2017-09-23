package main

import (
	"fmt"

	"github.com/abiosoft/ishell"
)

func stopCmd() *ishell.Cmd {
	return &ishell.Cmd{
		Name:      "stop",
		Help:      "stop [component] stops executing of the staq components or a specific component",
		Completer: Completer,
		Func: func(c *ishell.Context) {
			pattern := "*/*" // Default to all

			if len(c.Args) > 0 {
				pattern = c.Args[0]
			}

			for _, procfile := range AppStaq.Procfiles {
				for _, proc := range procfile.Procs {
					if Glob(pattern, fmt.Sprintf("%s/%s", procfile.Name, proc.Name)) {
						proc.Stop()

						// TODO Wrap this in a context with timeout and force stop if needed
						proc.Wait()
					}
				}
			}
		},
	}
}
