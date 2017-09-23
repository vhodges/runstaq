package main

import (
	"fmt"
	"time"

	"github.com/abiosoft/ishell"
)

func startCmd() *ishell.Cmd {
	return &ishell.Cmd{
		Name:      "start",
		Aliases:   []string{"run", "begin"},
		Help:      "start [component] executing the staq components or a specific component",
		Completer: Completer,
		Func: func(c *ishell.Context) {
			pattern := "*/*" // Default to all

			if len(c.Args) > 0 {
				pattern = c.Args[0]
			}

			for _, procfile := range AppStaq.Procfiles {
				for _, proc := range procfile.Procs {
					if Glob(pattern, fmt.Sprintf("%s/%s", procfile.Name, proc.Name)) {
						proc.Start(procfile, time.Duration(AppStaq.Delay)*time.Second)
					}
				}
			}
		},
	}
}
