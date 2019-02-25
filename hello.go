package main

import "fmt"

func main() {

	fmt.Println("Starting fizbuzz")

	fizbuzz()
}

func fizbuzz() {
	for i := 1; i < 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println(i, "FizBuzz")
		}

		if i%3 == 0 {
			fmt.Println(i, "Fiz")

		}
		if i%5 == 0 {
			fmt.Println(i, "Buzz")
		}
	}
}
