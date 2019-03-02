package main

import "fmt"

func fizbuzz() {
	x := 1

	fmt.Println("Starting fizbuzz", x)

	fizbuzz()
}

func runfizbuzz() {
	for i := 1; i < 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println(i, "FizBuzz")
			fmt.Printf("%T\n", i)
		}

		if i%3 == 0 {
			fmt.Println(i, "Fiz")

		}
		if i%5 == 0 {
			fmt.Println(i, "Buzz")
		}
	}
}
