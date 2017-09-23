package main

import (
	"os"
	"os/exec"
	"time"

	"github.com/abiosoft/ishell"
	"github.com/fatih/color"
)

type Proc struct {
	Name    string
	Command string
	Running bool
	Error   error

	cmd   *exec.Cmd
	shell *ishell.Shell
}

func (proc *Proc) Status() string {

	if proc.Running {
		return color.GreenString("running")
	}

	return color.MagentaString("not running")
}

func (proc *Proc) Start(procfile *Procfile, delay time.Duration) {

	if proc.Running {
		proc.shell.Printf("  %s\n    already running\n", proc.Name)
		return
	}

	proc.shell.Printf("  Starting %s... ", proc.Name)

	// TODO Windows support
	proc.cmd = exec.Command("/bin/sh", "-c", proc.Command)

	proc.cmd.Dir = procfile.Path
	proc.cmd.Stdout = procfile.Stdout
	proc.cmd.Stderr = procfile.Stderr

	proc.Error = proc.cmd.Start()
	if proc.Error != nil {
		proc.shell.Println(proc.Error)
		proc.Running = false
	} else {
		proc.Running = true
		go proc.Wait()
	}

	time.Sleep(delay)

	proc.shell.Printf("  Done\n")
}

func (proc *Proc) Stop() {
	if !proc.Running {
		proc.shell.Printf("  %s    already stopped\n", proc.Name)
		return
	}
	proc.shell.Printf("  Stopping %s...", proc.Name)

	proc.Error = proc.cmd.Process.Signal(os.Kill)
	if proc.Error != nil {
		proc.shell.Println(proc.Error)
	}

	proc.shell.Printf("  Done\n")

	proc.Running = false
}

func (proc *Proc) Wait() {
	proc.cmd.Wait()
	proc.Running = false
}
