package main

import (
	"encoding/binary"
	"flag"
	"fmt"
	"github.com/jsh/go_mutate"
	"io/ioutil"
	"os"
)

func main() {
	
	nbytePtr := flag.Int("nbyte", 0, "nbyte")
	maskPtr := flag.Uint64("mask", 0x00, "mask")

	flag.Parse()

	mask := make([]byte, 2)
	if n := binary.PutUvarint(mask, *maskPtr); n > 2 {
		fmt.Printf("bad mask %d\n", *maskPtr)
		os.Exit(1)
	}
	filelist := flag.Args()
	if len(filelist) != 2 {
		fmt.Println("bad file list", filelist)
		os.Exit(1)
	}

	program, err := ioutil.ReadFile(filelist[0])
	if err != nil {
		fmt.Printf("Can't read %s: %s\n", filelist[0], err)
		os.Exit(1)
	}
	err = mutate.Mask(program, *nbytePtr, mask[0], filelist[1])
	if err != nil {
		fmt.Printf("mutate failed: %s\n", err)
		os.Exit(1)
	}
}
