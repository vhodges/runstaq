package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/abiosoft/ishell"
)

type Procfile struct {
	Name    string
	Path    string
	Procs   []*Proc
	Error   error
	Running bool
	shell   *ishell.Shell
}

/*
web: bundle exec thin start
job: bundle exec rake jobs:work
*/

func NewProcfile(path string, shell *ishell.Shell) *Procfile {

	name := filepath.Base(path)

	procfile := &Procfile{Path: path, Name: name, shell: shell}

	file, err := os.Open(filepath.Join(path, "Procfile"))
	if err != nil {
		procfile.Error = err
		return procfile
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()
		words := strings.SplitAfterN(s, ":", 2)
		procfile.Procs = append(procfile.Procs, &Proc{Name: words[0], Command: words[1], shell: shell})
	}

	if err := scanner.Err(); err != nil {
		procfile.Error = err
	}

	return procfile
}

func (procfile *Procfile) Status() string {

	if procfile.Error != nil {
		return fmt.Sprintf("%s", procfile.Error)
	}

	return ""
}

func (procfile *Procfile) Start() {
	procfile.shell.Printf("Starting %s\n", procfile.Name)
	for _, proc := range procfile.Procs {
		proc.Start()
	}
	procfile.shell.Printf("Done\n")
}

func (procfile *Procfile) Stop() {
	procfile.shell.Printf("Stopping %s\n", procfile.Name)
	for _, proc := range procfile.Procs {
		proc.Stop()
	}
	procfile.shell.Printf("Done\n")
}
