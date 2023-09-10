package main

import (
	"fmt"
	"os"
)

func main() {
	tar := Tardigrade{}

	fmt.Println("Count: ", len(os.Args))
	fmt.Println("Last:", os.Args[len(os.Args)-1])

	tar.AddField(os.Args[1], os.Args[2], os.Args[3])
}
