package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

// takes doer(compiler/interpreter) and the source/pathToSource for execution and
// also a bool bin to flag if the doer generates seperate bin that needs manual execution
func runCmd(doer string,pathToSource string,bin  bool){
	var cmd,Rcmd *exec.Cmd
	if bin {
		cmd  = exec.Command(doer,pathToSource,"-o","/tmp/a.out")
		Rcmd = exec.Command("/tmp/a.out")
	}else  {
		cmd  = exec.Command(doer,pathToSource)
	}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error during processing " , err )
	}
	if bin && err == nil {
		Rcmd.Stdout = os.Stdout
		Rcmd.Stderr = os.Stderr
		Rerr := Rcmd.Run()
		if Rerr != nil {
			fmt.Println("Error during execution " , Rerr )
		}
	}
}

// checks extensions and calls doer function 
// basically a wrapper function
func do(extension string,pathToSource string) {
	switch extension { 
	case ".c":
		fmt.Println("gcc")
		runCmd("gcc",pathToSource,true)	
	case ".cpp":
		fmt.Println("g++")
		runCmd("g++",pathToSource,true)	
	case ".java":
		fmt.Println("javac and java")
		runCmd("javac",pathToSource,false)	
	case ".py":
		fmt.Println("Python")
		runCmd("python",pathToSource,false)	
	case ".go":
		fmt.Println("Golang")
		runCmd("go run",pathToSource,false)	
	default :
		fmt.Println("unsuported extension/language")
	}
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
