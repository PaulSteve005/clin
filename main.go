package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)
var tmpDir string = os.TempDir()

// takes doer(compiler/interpreter) and the source/pathToSource for execution and
// also a bool bin to flag if the doer generates seperate bin that needs manual execution
func runCmd(doer string,pathToSource string,isCompilable  bool,binName string){
	var cmd,Rcmd *exec.Cmd
	var binPath = filepath.Join(tmpDir,binName)
	if isCompilable {
		cmd  = exec.Command(doer,pathToSource,"-o",binPath)
		Rcmd = exec.Command(binPath)
	}else  {
		cmd  = exec.Command(doer,pathToSource)
	}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error during processing " , err )
	}
	if isCompilable && err == nil {
		Rcmd.Stdout = os.Stdout
		Rcmd.Stderr = os.Stderr
		Rcmd.Stdin = os.Stdin
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
		runCmd("gcc",pathToSource,true,"a.out")	
	case ".cpp":
		fmt.Println("g++")
		runCmd("g++",pathToSource,true,"a.out")	
	case ".java":
		fmt.Println("javac and java")
		runCmd("javac",pathToSource,false,"a.out")	
	case ".py":
		fmt.Println("Python")
		runCmd("python",pathToSource,false,"a.out")	
	case ".go":
		fmt.Println("Golang")
		runCmd("go run",pathToSource,false,"a.out")	
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
