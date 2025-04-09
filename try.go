package main 
import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

var tempDir string = os.TempDir()
var defaultBinPath = filepath.Join(tempDir,"a.out")
var binPath string
var terminate bool

func main() {
	// define a flag: -o
	// Parse the command-line flags

	// Remaining arguments (non-flags) are source files


	

	flag.StringVar(&binPath, "o",defaultBinPath, "path to the bin file")
	flag.BoolVar(&terminate, "t", false, "tells clin weather to run the bin or not")
	flag.Parse()

	if terminate {
		fmt.Println("terminate activated")
	}
	if binPath == defaultBinPath {
		fmt.Println("default bin")
	}else{
		fmt.Println(binPath)
	}
	buildFlags := flag.Args()

	if len(buildFlags) < 1 {
		fmt.Println("Please specify a source file.")
		os.Exit(1)
	}
	fmt.Println("Source File:", buildFlags)
	fmt.Println("Output File:", binPath)
}

