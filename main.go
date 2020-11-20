package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Enter listen port as command agrument")
		return
	}
	fmt.Println(os.Args[1])
}
