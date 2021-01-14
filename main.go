package main

import (
	"fmt"
)

func main() {
	port := 3000
	_, err := startWebServer(port)
	fmt.Println(err)
}

func startWebServer(port int) (int, error) {
	fmt.Println("Starting server...")
	fmt.Println("Server started on port", port)
	return port, nil
}
