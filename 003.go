package main

import (
	"fmt"
	"time"
	"os"
	"net"
)

func main() {
	fmt.Println("Welcome to the playground!!")

	fmt.Println("The time is", time.Now())

	fmt.Println("And if you try to open a file:")
	fmt.Println(os.Open("003.go"))  // succeeds here :D

	fmt.Println("Or access the network:")
	fmt.Println(net.Dial("tcp", "google.com"))
}
