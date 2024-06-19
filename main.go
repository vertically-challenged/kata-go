package main

import (
	"fmt"

	"calc/src/controller"
)

func main() {
	for {
		fmt.Println("Input:")
		var exception string
		fmt.Scan(&exception)
		fmt.Println("Output:")
		fmt.Println(controller.Controller(exception))
	}
}
