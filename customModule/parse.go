package customModule 

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)
func isSupported(extension string)bool {
	switch extension {
	case ".c", ".cpp", ".py", ".go", ".zig",".rs",".java",".swift":
		return true
	default:
		return false
	}
}


//  variables //
var ( 
	FoundBin = false 
	FoundSource = false 
	NoExecBin = false 

	PathToSource = ""   // stores path to source code
	BinPath = ""      // stores path to binary blob
	BuildFlags = ""
)

const version = "alpha0.3.1"
const helpText = `
Usage: clin [options] [build flags / options] <source-file>

Options:
-v or --version    print version
-h or --help       Show this help message

-o <file>          Set the output binary file path
-ot <file>         Set path and not run the binary after building
--build            Everything after this is considered build flags

This is not necessary but is there if u need it
	
Examples:
clin -o bin/myapp test.c
clin -ot bin/myapp hello.cpp
clin script.py
clin --build "
clin -u myscript.py --input data.txt --verbose -n 5
clin myscript.py --build -u--input data.txt --verbose -n 5
`

// end variables //

func ParseArgs(parameters[]string){
	//parameters
	fmt.Println(parameters)

	inBuildMode := false


	for i := 0; i<len(parameters);i++ {
		var word string = parameters[i]

		if inBuildMode {    // if --build is encountered  ignores all parsing and store everything in buildFlags
			BuildFlags += word + " "
			continue
		}

		if strings.HasPrefix(word, "--") {
			switch word {
			case "--help":
				fmt.Print(helpText)
				os.Exit(0)
			case "--version":
				fmt.Println(version)
				os.Exit(0)
			case "--build":
				if FoundSource {
					inBuildMode = true
					continue	
				} else {
					fmt.Println("No source files found")
					os.Exit(2)
				}
				break
			default:
				// Unknown flags as buildFlags
				BuildFlags += word + " "
			}
		} else if strings.HasPrefix(word, "-") {
			switch word {
			case "-o":
				FoundBin = true
			case "-ot":
				FoundBin = true
				NoExecBin = true
			case "-h":
				fmt.Print(helpText)
				os.Exit(0)
			case "-v":
				fmt.Println(version)
				os.Exit(0)
			default:
				// Unknown flags as buildFlags
				BuildFlags += word + " "
			}
		
	} else if isSupported(filepath.Ext(word)) {
		if !FoundSource {
			FoundSource = true
			PathToSource = word
		} else {
			fmt.Println("Two source Found !!")
			os.Exit(2)
		}
	} else if FoundBin && BinPath == "" {
		BinPath = word
	} else {
		BuildFlags += word + " "
	}
	}


	if !FoundSource {
		fmt.Println("No valid source file found or unsupported language.")
		fmt.Println("Captured build flags:", BuildFlags)
		os.Exit(2)
	}
	// In parse.go, at the end of ParseArgs:
	BuildFlags = strings.TrimSpace(BuildFlags)


}

