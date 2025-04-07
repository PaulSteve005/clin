package main

import (
	"fmt"
	"os"
)
func exec(extension string) {
	switch extension { 
	case "c":
		fmt.Println("gcc")
	case "cpp":
		fmt.Println("g++")
	case "java":
		fmt.Println("javac and java")
	case "py":
		fmt.Println("Python")
	case "go":
		fmt.Println("Golang")
	case "js":
		fmt.Println("node-js")
	}
}

func main()  {
	//initial sanity check
	if  len(os.Args) <   2{
		fmt.Println("no source code specified")
	}
	filePath :=  os.Args[1]
	fmt.Println("path entered ", filePath)
	exec(filePath)

}
