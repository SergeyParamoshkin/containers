package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	fmt.Println(os.Getenv("BLABLA"))
	fmt.Println("Hello rebrain!")
	run()
}

func run() {
	fmt.Println(runtime.NumCPU())
}
