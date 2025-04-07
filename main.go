package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)
func do(extension string,pathToSource string) {
	var cmd *exec.Cmd 
	switch extension { 
	case ".c":
		fmt.Println("gcc")
		cmd = exec.Command("gcc" ,pathToSource,"-o","/tmp/a.out")
	case ".cpp":
		fmt.Println("g++")
	case ".java":
		fmt.Println("javac and java")
	case ".py":
		fmt.Println("Python")
	case ".go":
		fmt.Println("Golang")
	case ".js":
		fmt.Println("node-js")
	}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println("error during processing " , err )
	}
	run := exec.Command("/tmp/a.out")
	run.Stdout = os.Stdout
	run.Stderr = os.Stderr
	run.Run()
}

func main()  {
	//initial sanity check
	if  len(os.Args) <   2{
		fmt.Println("no source code specified")
		return
	}
	pathToSource := os.Args[1]
	fmt.Println("path entered ", pathToSource)

	// extension extraction
	extension :=  filepath.Ext(pathToSource)
	if extension != ""{
		do(extension,pathToSource)
	}else{
		fmt.Println("no extension found ")
	}
	return
}
