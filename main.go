package main

import (
	"fmt"
	"go-start/testpackage"
)

func main() {
	fmt.Println("Hello, World!")

	var name string = "Jone"



	testpackage.Test(name)
}
