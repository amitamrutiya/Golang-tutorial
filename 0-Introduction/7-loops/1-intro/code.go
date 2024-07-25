package main

import (
	"fmt"
)

func bulkSend(numMessages int) float64 {
	var cost float64
	for i := 0; i < numMessages; i++ {
		// j := float64(i / 100)
		cost = cost + 1.00 + (float64(i) * 0.01)
	}
	return cost
}

// don't edit below this line

func test(numMessages int) {
	fmt.Printf("Sending %v messages\n", numMessages)
	cost := bulkSend(numMessages)
	fmt.Printf("Bulk send complete! Cost = %f\n", cost)
	fmt.Println("===============================================================")
}

func main() {
	test(10)
	test(20)
	test(30)
	test(40)
	test(50)
}
