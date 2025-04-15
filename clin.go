package main

import (
	"clin/customModule"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)
var tmpDir string = os.TempDir()
var isWierd bool = false // specially for zig
var seperator string = "/"  //initialized it to posix std

// takes doer(compiler/interpreter) and the source/pathToSource for execution and much more actually
// all function parameters when set to "" means get default value xcept for doer,pathToSource(needed anyhow),bools  and runner's "" means no runner required
// " " is for none for all the parameters xcept the exceptions mentioned above
func runCmd(doer string,doer2 string, pathToSource string, isCompilable bool, binName string, binoption string, runner string) {
	var cmd, Rcmd *exec.Cmd //cmd for compile/interpreter command and Rcmd for runner command
	var binPath = ""
	var args []string

	// binoptions //
	if binoption == ""{
		binoption = "-o"  // setting default binoption
	}

	// binary path //
	if binName == " " {    // i.e. no bin requiered
		binPath = binName
	} else if customModule.BinPath == "" && binName != "" {
		binPath = filepath.Join(tmpDir, binName)  // default bin
	} else {
		binPath = customModule.BinPath   //custom bin user defined !!
	}

	// Verbose //
	customModule.LogVerbose("BinOption    : %s", binoption)
	customModule.LogVerbose("BinPath      : %s", binPath)

	// checking if doer is installed or not //
	if resolvedPath, err := exec.LookPath(doer); err == nil {
		customModule.LogVerbose("Executable   : %s\n", resolvedPath)
	} else {
		fmt.Printf("Executable   : not found (%v) ", err)
	}



	// support for 2 doers //
	if doer2 != " "{
		args = append(args,doer2)
	}

	// binoptions  defaults to /tmp //
	if isCompilable && binoption != " " {  // ignores stuff  if binoptions is set to " " none
		if isWierd {
			args = append(args, binoption+binPath)
		} else {
			args = append(args,binoption, binPath)
		}
	}

	// build flags //
	if customModule.BuildFlags != "" {
		args = append(args, customModule.BuildFlags)  // append build flags only if there exists  build flags
	}

	args = append(args, pathToSource) // source is necessary !!

	// Compiler / Interpreter //
	customModule.LogVerbose("Compilation Command : %s %v", doer, args)
	cmd = exec.Command(doer,args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin  //stdin cuse interpreter is here
	Ctime := time.Now()  //start clock
	doErr := cmd.Run()
	elapsed := time.Since(Ctime)  // calculate time
	if doErr != nil {
		fmt.Println("Error during processing ", doErr)
	} else  {
		customModule.LogVerbose("Compiled/Interpreted Successfll :)) %s\n",elapsed)
	}

	// Runner //
	if isCompilable && doErr == nil && !customModule.NoExecBin {
		if runner == ""{                                                                    // normal	
			if !strings.Contains(binPath,seperator) {
				binPath = "."+seperator+binPath

			}
			customModule.LogVerbose("Runtime Command     : %s ", binPath)
			Rcmd = exec.Command(binPath) // Execute the generated binary directly
		} else if runner == "java" {													//java things

			customModule.LogVerbose("Processing things for java ")
			sourceName := filepath.Base(pathToSource[:len(pathToSource)-len(filepath.Ext(pathToSource))])
			directory := pathToSource[:len(pathToSource)-len(sourceName)-5] 

			if directory != ""{                                                        // if source was in a directory
				//customModule.LogVerbose(" Directory where class is located %s ",directory)
				//customModule.LogVerbose(" Name of the class file %s ",sourceName)

				customModule.LogVerbose("Runtime Command     : %s -cp %s %s", runner, directory, sourceName)
				Rcmd = exec.Command(runner,"-cp",directory,sourceName)
			}else {
				//customModule.LogVerbose(" Name of the class file %s ",sourceName)
				customModule.LogVerbose("Runtime Command     : %s %s", runner, sourceName)
				Rcmd = exec.Command(runner,sourceName )

			}
		} else {     // other runners (for future use)
			Rcmd = exec.Command(runner,binPath)
		}

		Rcmd.Stdout = os.Stdout
		Rcmd.Stderr = os.Stderr
		Rcmd.Stdin = os.Stdin
		Ctime = time.Now()  //start clock
		Rerr := Rcmd.Run()
		elapsed = time.Since(Ctime)  // calculate time
		if Rerr != nil {
			fmt.Println("Error during execution ", Rerr)
		}else {
			customModule.LogVerbose("Running Successfll :)) %s",elapsed)
		}
	}
}



// checks extensions and calls doer function 
// basically a wrapper function
func do() {
	extension :=  filepath.Ext(customModule.PathToSource)
	// proper  os specific extension
	var Defaultbin string = "a.out"
	if runtime.GOOS == "windows"{  // windows things
		Defaultbin = "a.exe"
		seperator = "\\"  
	}


	switch extension { 
	case ".c":
		customModule.LogVerbose("Detected C ")
		runCmd("gcc"," ", customModule.PathToSource, true, Defaultbin,"", "")	
	case ".cpp":
		customModule.LogVerbose("Detected C++ ")
		runCmd("g++"," ", customModule.PathToSource, true,Defaultbin, "", "")	
	case ".go":
		customModule.LogVerbose("Detected Golang ")
		runCmd("go","build",customModule.PathToSource,true,Defaultbin,"", "")	
	case ".rs":
		customModule.LogVerbose("Detected Rust ")
		runCmd("rustc"," ", customModule.PathToSource, true,Defaultbin, "", "")
	case ".swift":
		customModule.LogVerbose("Detected Swift ")
		runCmd("swiftc"," ", customModule.PathToSource, true,Defaultbin, "", "")
	case ".zig":
		customModule.LogVerbose("Detected Zig ")
		isWierd = true
		runCmd("zig","build-exe",customModule.PathToSource,true,Defaultbin,"-femit-bin=","")
	case ".f90",".f95",".f",".f03",".f08",".for":
		customModule.LogVerbose("Detected Fortran ")
		runCmd("gfortran"," ",customModule.PathToSource,true,Defaultbin, "", "")
	case ".hs":
		customModule.LogVerbose("Detected Haskel")
		runCmd("ghc"," ",customModule.PathToSource,true,Defaultbin, "", "")

	// mixed  //	
	case ".java":
		customModule.LogVerbose("Detected Java")
		if customModule.FoundBin{
			fmt.Println("java does not support customizable directory")
			os.Exit(3)
		}else {
			runCmd("javac"," ",customModule.PathToSource,true," "," ","java")	
		}

	// Interpreted //
	case ".py":
		customModule.LogVerbose("Detected Python ")
		runCmd("python"," ",customModule.PathToSource,false,"", "", "")	
	case ".rb":
		customModule.LogVerbose("Detected Ruby ")
		runCmd("ruby"," ",customModule.PathToSource,false,"", "", "")	
	case ".pl":
		customModule.LogVerbose("Detected Perl")
		runCmd("perl"," ",customModule.PathToSource,false,"", "", "")	
	case ".js":
		customModule.LogVerbose("Detected Javascript ")
		runCmd("node"," ",customModule.PathToSource,false,"", "", "")
	case ".php":
		customModule.LogVerbose("Detected PHP ")
		runCmd("php"," ",customModule.PathToSource,false,"", "", "")
	case ".lua":
		customModule.LogVerbose("Detected lua ")
		runCmd("lua"," ",customModule.PathToSource,false,"", "", "")
	case ".ts":       //  still interpreted
		customModule.LogVerbose("Detected TypeScript ")
		sourceName := filepath.Base(customModule.PathToSource[:len(customModule.PathToSource)-len(filepath.Ext(customModule.PathToSource))])
		runCmd("tsc"," ",customModule.PathToSource,true,sourceName+".js", "--outFile", "node")
	default :
		fmt.Println("unsuported extension/language ")  // don't think it ll get here but just incase
		fmt.Println("Soeething is not right with clin you shouldn't see this message ")
		os.Exit(69)  //:P
	}
}

func main()  {
	// parse args  //
	customModule.ParseArgs(os.Args[1:])

	// do the thing //
	do()
}
