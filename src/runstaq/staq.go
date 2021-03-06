package main

import (
	"github.com/abiosoft/ishell"
)

type Staq struct {
	Procfiles []*Procfile
	Delay     int64 // How to wait between process invocations

	shell *ishell.Shell
}

func buildStaq(paths []string, shell *ishell.Shell) *Staq {

	staq := &Staq{shell: shell}

	for _, path := range paths {
		staq.Procfiles = append(staq.Procfiles, NewProcfile(path, shell))
	}

	return staq
}

func (staq *Staq) module(name string) *Procfile {
	var p *Procfile = nil

	for _, procfile := range staq.Procfiles {
		if procfile.Name == name {
			p = procfile
			break
		}
	}

	return p
}

func (staq *Staq) shutdown() {
	staq.shell.Printf("Shutting down...\n")
	for _, procfile := range AppStaq.Procfiles {
		for _, proc := range procfile.Procs {

			proc.Stop()

			// TODO Figure out wating for procs to stop (before it was on the procfile as a set
			// but since we're being more flexible now, so does the code for that.

			// TODO Wrap this in a context with timeout and force kill them
			//procfile.wg.Wait() // Wait for them to all terminate
		}
	}
}
