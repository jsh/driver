package main

import (
	"fmt"
	"github.com/jsh/go_tryprog"
	"io/ioutil"
	"os"
)

func main() {

	args := os.Args[1:]
	fmt.Println("file list", args)
	if len(args) != 2 {
		fmt.Println("bad file list", args)
		os.Exit(1)
	}

	expected, err := ioutil.ReadFile(args[0])
	if err != nil {
		fmt.Println("can't read ", args[0])
		os.Exit(1)
	}
	
	err = tryprog.Try(expected, args[1])
	fmt.Println("result is", err)
}
