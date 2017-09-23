package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

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
		procfile.Procs = append(procfile.Procs, &Proc{Name: strings.TrimSuffix(words[0], ":"), Command: words[1], shell: shell})
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
