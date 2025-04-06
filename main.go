package main

import (
	"fmt"
	"os"
)

func main()  {
	//initial sanity check
	if  len(os.Args) <   2{
		fmt.Println("no source code specified")
	}
	var filePath string =  os.Args[1]
	fmt.Println("path entered ", filePath)
}
