// this function takes a argument of any file type and prints it out in the cmd terminal
// cmd = "run main.go myfile.txt"
//
// EXTRA
// to use the code to print out the code itself, one must first compile the code and later use:
// cmd = "main.exe main.go"
package main

import (
	"fmt"
	"os"
)

func main() {
	// fmt.Println(os.Args[1])
	txt, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(string(txt))
}
