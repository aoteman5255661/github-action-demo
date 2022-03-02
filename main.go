package main

import "fmt"

func Cat() string {
	return "wanging"
}

func main() {
	saying := Cat()
	fmt.Print(saying)
}
