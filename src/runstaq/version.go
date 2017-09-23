package main

import (
	"github.com/abiosoft/ishell"
)

var release = "2017.10"
var build_sha string

func versionCmd() *ishell.Cmd {
	// register a function for "greet" command.
	return &ishell.Cmd{
		Name: "version",
		Help: "print runstaq version/build info",
		Func: func(c *ishell.Context) {
			c.Printf("runstaq version: %s SHA: %s\n", release, build_sha)
		},
	}
}
