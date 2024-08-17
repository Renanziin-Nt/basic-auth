package main

import (
	"User_Service/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	r := router.RoutersGenerator()
	fmt.Printf("Starting server on: %d", 8080)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", 8080), r))
}
