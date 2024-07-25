package main

import "fmt"

func main() {
	const name = "Saul Goodman"
	const openRate = 30.5

	var msg = fmt.Sprintf("Hi %s your open rate is %f percent", name, openRate)

	// don't edit below this line

	fmt.Println(msg)
}
