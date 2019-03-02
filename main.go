package main

import "fmt"

type xireon int

var a xireon
var b int

func main() {
	a = 1
	b = int(a)
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println(b)
	fmt.Printf("%T\n", b)
}
