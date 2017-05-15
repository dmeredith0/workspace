package main

import "fmt"

func main() {
	//var x string
	//This is a comment x=world()
	fmt.Printf("Hello %s or whatever you like\n", world("stringhowdy"))
}

func world(x string) string {
	return x
}
