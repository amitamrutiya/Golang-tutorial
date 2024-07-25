package main

import (
	"fmt"
	"log"
	"net/http"
	"mongodb/router"
)

func main() {
	fmt.Println("Hello, World!")
	r := router.Router()
	fmt.Println("Server is getting Strated...")
	log.Fatal(http.ListenAndServe(":8080", r))
	fmt.Println("Server is running on port 8080...")
}
