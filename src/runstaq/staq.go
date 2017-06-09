package main

import (
	"github.com/abiosoft/ishell"
)

type Staq struct {
	Procfiles []*Procfile
	shell     *ishell.Shell
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
		procfile.Stop()
	}
}
