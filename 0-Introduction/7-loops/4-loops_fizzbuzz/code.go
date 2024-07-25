package main

import "fmt"

func fizzbuzz() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Println("fizz")
		} else if i%5 == 0 {
			fmt.Println("buzz")
		} else if i%3 == 0 && i%5 == 0 {
			fmt.Println("fizzbuzz")
		}
	}
}

// don't touch below this line

func main() {
	fizzbuzz()
}
