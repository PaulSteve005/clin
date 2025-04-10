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


// takes doer(compiler/interpreter) and the source/pathToSource for execution and
// also a bool bin to flag if the doer generates seperate bin that needs manual execution
func runCmd(doer string, pathToSource string, isCompilable bool, binName string, binoption string, runner string) {
    var cmd, Rcmd *exec.Cmd
    var binPath = "  "

    // binary path //
    if customModule.BinPath == "" && binName != "" {
        binPath = filepath.Join(tmpDir, binName)
    } else {
        binPath = customModule.BinPath
    }

    //debug
    fmt.Println("BinOption:", binoption)
    fmt.Println("BinPath:", binPath)
    fmt.Println("PATH:", os.Getenv("PATH"))
    fmt.Println(exec.LookPath(doer))

    var args []string
    args = append(args, customModule.BuildFlags)
    args = append(args, pathToSource)

    if isCompilable && customModule.FoundBin {
        args = append(args, "-o", binPath)
    }

    cmd = exec.Command(doer, args...)

    // Compiler / Interpreter //
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    cmd.Stdin = os.Stdin
    err := cmd.Run()
    if err != nil {
        fmt.Println("Error during processing ", err)
    }

    // Runner //
    if isCompilable && err == nil && !customModule.NoExecBin && customModule.FoundBin {
        Rcmd = exec.Command(binPath) // Execute the generated binary directly
        Rcmd.Stdout = os.Stdout
        Rcmd.Stderr = os.Stderr
        Rcmd.Stdin = os.Stdin
        Rerr := Rcmd.Run()
        if Rerr != nil {
            fmt.Println("Error during execution ", Rerr)
        }
    } else if isCompilable && err == nil && !customModule.NoExecBin && runner != "" { // For cases like Java
        Rcmd = exec.Command(runner, filepath.Base(customModule.PathToSource[:len(customModule.PathToSource)-len(filepath.Ext(customModule.PathToSource))])) // Extract class name
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
        binOptionToPass := ""
        if customModule.FoundBin {
            binOptionToPass = "-o "
        }
        runCmd("gcc", customModule.PathToSource, true, Defaultbin, binOptionToPass, "")	
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
			runCmd("javac",customModule.PathToSource,true,""," ","java")	
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
