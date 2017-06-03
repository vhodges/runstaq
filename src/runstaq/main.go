package main

import (
	//	"strings"
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/abiosoft/ishell"
)

var AppStaq *Staq
var Names []string

var Completer = func([]string) []string {
	return Names
}

func main() {

	if len(os.Args[1:]) < 1 {
		fmt.Printf("Usage: runstaq Folder1 Folder2 FolderN\n")
		fmt.Printf("   or: runstaq -stack stackfile Folder1 FolderN\n")
		fmt.Printf("   or: runstaq -stack stackfile\n")
		os.Exit(1)
	}

	staqFile := flag.String("stack", "", "filename of file containing list of folders in the stack, one per line.")
	flag.Parse()

	var modules = make([]string, 0, 20)

	if *staqFile != "" {
		modules = readStaqFile(modules, *staqFile)
	}

	for _, p := range flag.Args() {
		modules = append(modules, p)
		b := filepath.Base(p)
		Names = append(Names, b)
	}

	shell := ishell.New()

	shell.SetHomeHistoryPath(".runstaq_history")

	shell.Printf("Welcome to runstaq!!\n\nType help to get started\n\n")

	shell.AddCmd(versionCmd())
	shell.AddCmd(listCmd())
	shell.AddCmd(startCmd())
	shell.AddCmd(stopCmd())

	AppStaq = buildStaq(modules, shell)

	// run shell
	shell.Run()
}

func readStaqFile(modules []string, filename string) []string {

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()
		modules = append(modules, s)
		Names = append(Names, filepath.Base(s))
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return modules
}
