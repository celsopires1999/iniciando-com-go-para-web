package main

import "fmt"

func main() {

	celso := "Celso"
	fmt.Println("P1")
	// t := name()
	fmt.Println("P2")
	if t := name(); t == celso {
		fmt.Println("=")
	} else {
		fmt.Println("!=")
	}

	fmt.Println("P3")
}

func name() string {
	return "Celso"
}
