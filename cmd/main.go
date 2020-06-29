package main

import (
	"os"
	"log"
	"fmt"

	"github.com/LumaKernel/brainhug"
)


const usage = `BrainHug - A Brain*uck implementation written in Go
Usage: brainhug <program> <input>
`

func printUsage() {
	fmt.Print(usage)
}

func main() {
	if len(os.Args) != 2 && len(os.Args) != 3 {
		printUsage()
		return
	}
	program := os.Args[1]
	input := []byte{}

	if len(os.Args) == 3 {
		input = []byte(os.Args[2])
	}

	stdout, err := brainhug.Proceed(program, []byte(input))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(string(stdout))
}
