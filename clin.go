package main

import (
	"clin/customModule"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	//"strings"
)
var tmpDir string = os.TempDir()
var isWierd bool = false // specially for zig

// takes doer(compiler/interpreter) and the source/pathToSource for execution and
// also a bool bin to flag if the doer generates seperate bin that needs manual execution + binoption biname runner and support for -ot flag
func runCmd(doer string, pathToSource string, isCompilable bool, binName string, binoption string, runner string) {
	var cmd, Rcmd *exec.Cmd
	var binPath = ""
	var args []string

	// binoptions //
	if binoption == ""{
		binoption = "-o"  // setting default binoption
	}

	// binary path //
	if binName == "" {    // i.e. no bin requiered
		binPath = binName
	} else if customModule.BinPath == "" && binName != "" {
		binPath = filepath.Join(tmpDir, binName)  // default bin
	} else {
		binPath = customModule.BinPath   //custom bin user defined !!
	}

	//debug
	fmt.Println("BinOption:", binoption)
	fmt.Println("BinPath:", binPath)
	fmt.Println("PATH:", os.Getenv("PATH"))
	fmt.Println(exec.LookPath(doer))

	if customModule.BuildFlags != "" {
		args = append(args, customModule.BuildFlags)  // append build flags only if there exists  build flags
	}
	args = append(args, pathToSource) // source is necessary !!

	if isCompilable && binoption != " " {  // ignores stuff  if binoptions is set to " " none
		if isWierd {
			args = append(args, binoption+binPath)
		} else {
			args = append(args,binoption, binPath)
		}
	}

	// Compiler / Interpreter //
	cmd = exec.Command(doer, args...)
	fmt.Println("compilation command ",doer,args)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin  //stdin cuse interpreter is here
	doErr := cmd.Run()
	if doErr != nil {
		fmt.Println("Error during processing ", doErr)
	}

	// Runner //
	if isCompilable && doErr == nil && !customModule.NoExecBin {
		if runner == ""{                                                                    // normal
			Rcmd = exec.Command(binPath) // Execute the generated binary directly
		} else if runner == "java" {																												//java things
			sourceName := filepath.Base(pathToSource[:len(pathToSource)-len(filepath.Ext(pathToSource))])
			directory := pathToSource[:len(pathToSource)-len(sourceName)-5] 
				fmt.Println("dir",directory,"\n source  name ",sourceName)
			if directory != ""{
				Rcmd = exec.Command(runner,"-cp",directory,sourceName)
				fmt.Println("runner command ",runner,"-cp",directory,sourceName)
			}else {
				Rcmd = exec.Command(runner,sourceName )
				fmt.Println("runner command ",runner,sourceName)
			}
		} else {																																	// other runners (for future use)
			Rcmd = exec.Command(runner,binPath)
		}

		Rcmd.Stdout = os.Stdout
		Rcmd.Stderr = os.Stderr
		Rcmd.Stdin = os.Stdin
		Rerr := Rcmd.Run()
		if Rerr != nil {
			fmt.Println("Error during execution ", Rerr)
		}
	}
}



// checks extensions and calls doer function 
// basically a wrapper function
func do() {
	extension :=  filepath.Ext(customModule.PathToSource)
	// proper  os specific extension
	var Defaultbin string = "a.out"
	if runtime.GOOS == "windows"{
		Defaultbin = "a.exe"
	}


	switch extension { 
	case ".c":
		fmt.Println("gcc")
		runCmd("gcc", customModule.PathToSource, true, Defaultbin,"", "")	
	case ".cpp":
		fmt.Println("g++")
		runCmd("g++", customModule.PathToSource, true,Defaultbin, "", "")	
	case ".go":
		fmt.Println("Golang")
		runCmd("go build",customModule.PathToSource,true,Defaultbin,"", "")	
	case ".rs":
		fmt.Println("rustc")
		runCmd("rustc", customModule.PathToSource, true,Defaultbin, "", "")
	case ".swift":
		fmt.Println("swiftc")
		runCmd("swiftc", customModule.PathToSource, true,Defaultbin, "", "")
	case ".zig":
		fmt.Println("Zig")
		runCmd("zig build-exe",customModule.PathToSource,true,Defaultbin,"-femit-bin=","")
		isWierd = true
	case ".f90",".f95",".f",".f03",".f08",".for":
		runCmd("gfortran",customModule.PathToSource,true,Defaultbin, "", "")
	case ".hs":
		runCmd("ghc",customModule.PathToSource,true,Defaultbin, "", "")

	// mixed  //	
	case ".java":
		fmt.Println("javac and java")
		if customModule.FoundBin{
			fmt.Println("java does not support customizable directory")
			os.Exit(3)
		}else {
			runCmd("javac",customModule.PathToSource,true," "," ","java")	
		}

	// Interpreted //
	case ".py":
		fmt.Println("Python")
		runCmd("python",customModule.PathToSource,false,"", "", "")	
	default :
		fmt.Println("unsuported extension/language")
	}
}

func main()  {
	//initial sanity check
	customModule.ParseArgs(os.Args[1:])

	fmt.Println("path entered ", customModule.PathToSource)

	// extension extraction
	do()
}
