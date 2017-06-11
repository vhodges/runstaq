package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/abiosoft/ishell"
)

type Procfile struct {
	Name  string
	Path  string
	Procs []*Proc
	Error error

	Stdout io.Writer
	Stderr io.Writer

	shell *ishell.Shell
	wg    sync.WaitGroup
}

/*
web: bundle exec thin start
job: bundle exec rake jobs:work
*/

// TODO Clean this up
func NewProcfile(path string, shell *ishell.Shell) *Procfile {

	name := filepath.Base(path)

	procfile := &Procfile{Path: path, Name: name, shell: shell}

	var err error
	var stdout *os.File
	var stderr *os.File
	var file *os.File

	file, err = os.Open(filepath.Join(path, "Procfile"))
	if err != nil {
		procfile.Error = err
		shell.Println(err)
		return procfile
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()
		words := strings.SplitAfterN(s, ":", 2)
		procfile.Procs = append(procfile.Procs, &Proc{Name: words[0], Command: words[1], shell: shell})
	}

	err = scanner.Err()
	if err != nil {
		procfile.Error = err
		shell.Println(err)
	}

	// stdout
	stdout, err = os.Create(filepath.Join(path, "runstaq.stdout.log"))
	if err != nil {
		procfile.Error = err
		shell.Println(err)
		return procfile
	}
	procfile.Stdout = stdout

	stderr, err = os.Create(filepath.Join(path, "runstaq.stderr.log"))
	if err != nil {
		procfile.Error = err
		shell.Println(err)
		return procfile
	}
	procfile.Stderr = stderr

	return procfile
}

func (procfile *Procfile) Status() string {

	if procfile.Error != nil {
		return fmt.Sprintf("%s", procfile.Error)
	}

	return ""
}

func (procfile *Procfile) Start(delay int64) {

	procfile.shell.Printf("%s\n", procfile.Name)

	for _, proc := range procfile.Procs {
		proc.Start(procfile, time.Duration(delay)*time.Second)
	}
	procfile.shell.Printf("\n")
}

func (procfile *Procfile) Stop() {

	procfile.shell.Printf("%s\n", procfile.Name)

	for _, proc := range procfile.Procs {
		proc.Stop()
	}

	// TODO Wrap this in a context with timeout and force kill them
	procfile.wg.Wait() // Wait for them to all terminate

	procfile.shell.Printf("\n")
}
