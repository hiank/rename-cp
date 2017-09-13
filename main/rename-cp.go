package main

import (
	"fmt"
	"flag"
	"github.com/hiank/rename-cp"
)

func main() {

	srcDir := flag.String("s", "", "src dir")
	dstDir := flag.String("d", "", "dst dir")
	// mixLen := flag.Int("l", 300, "the num of the mix byte")

	flag.Parse()

	switch {
	case *srcDir == "": 
		fmt.Println("should define srcDir with -s") 
		fallthrough
	case *dstDir == "": 
		fmt.Println("should define dstDir with -d") 
		return
	}
	rc.DuplicateDir(*srcDir, *dstDir)
}