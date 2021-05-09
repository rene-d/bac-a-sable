package main

import (
	"fmt"
	"runtime/debug"
)

func main() {
	fmt.Println("hello")
	if info, available := debug.ReadBuildInfo(); available {
		fmt.Println(info.Main.Version)
	}
}
