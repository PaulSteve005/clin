package main

import (
	"fmt"
	"os"
	"strings"
	"path/filepath"
)
func isSupported(extension string)bool {
	switch extension {
	case ".c", ".cpp", ".py", ".go":
		return true
	default:
		return false
	}
}


//  variables //
var foundBin bool = false 
var foundSource bool = false 
var noExecBin bool = false 

var sourcePath string = ""   // stores path to source code
var binPath string = ""      // stores path to binary blob
var buildFlags string = ""

const version = "alpha0.3.1"
const helpText = `
Usage: clin [options] [build flags / options] <source-file>

Options:
-v or --version    print version
-h or --help       Show this help message

-o <file>          Set the output binary file path
-ot <file>         Set path and not run the binary after building
	
Examples:
clin -o bin/myapp test.c
clin -ot bin/myapp hello.cpp
clin script.py
clin --build "
clin -u myscript.py --input data.txt --verbose -n 5
`

// end variables //

func main(){
	//parameters
	parameters := os.Args[1:]
	fmt.Println(parameters)	


	for i := 0; i<len(parameters);i++ {
		var word string = parameters[i]
		if strings.HasPrefix(word, "--") {
			switch word {
			case "--help":
				fmt.Println(helpText)
				os.Exit(0)
			case "--version":
				fmt.Println(version)
				os.Exit(0)
			default:
				// Unknown flags as buildFlags
				buildFlags += word + " "
			}
		} else if strings.HasPrefix(word, "-") {
			switch word {
			case "-o":
				foundBin = true
			case "-ot":
				foundBin = true
				noExecBin = true
			case "-h":
				fmt.Println(helpText)
				os.Exit(0)
			case "-v":
				fmt.Println(version)
				os.Exit(0)
			default:
				// Unknown flags as buildFlags
				buildFlags += word + " "
			}
		
	} else if isSupported(filepath.Ext(word)) {
		foundSource = true
		sourcePath = word
	} else if foundBin && binPath == "" {
		binPath = word
	} else {
		buildFlags += word + " "
	}
	}


	if !foundSource {
		fmt.Println("No valid source file found or unsupported language.")
		fmt.Println("Captured build flags:", buildFlags)
		os.Exit(2)
	}


fmt.Println(sourcePath,"\n", binPath, "\n", buildFlags, "\n")
fmt.Println(foundBin,"\n" ,foundSource, "\n" ,noExecBin, "\n")
}

