package main

import (
	"os/exec"

	"github.com/abiosoft/ishell"
)

type Proc struct {
	Name    string
	Command string
	Running bool

	cmd   *exec.Cmd
	shell *ishell.Shell
}

func (proc *Proc) Status() string {

	if proc.Running {
		return "running"
	}

	return "not running"
}

func (proc *Proc) Start() {

	if proc.Running {
		proc.shell.Printf("%s already running\n", proc.Name)
		return
	}

	proc.shell.Printf("  Starting %s...\n", proc.Name)
	proc.shell.Printf("    Exec: %s ", proc.Command)
	proc.shell.Printf("  Done\n")

	proc.Running = true
}

func (proc *Proc) Stop() {
	if !proc.Running {
		proc.shell.Printf("%s already stopped\n", proc.Name)
		return
	}
	proc.shell.Printf("  Stopping %s...", proc.Name)
	proc.shell.Printf("  Done\n")

	proc.Running = false
}
