package main

import (
	"assignment4/internal/constants"
	"assignment4/internal/server"

	"fmt"
)

func main() {
	fmt.Println("Sample Assignment")
	server := server.NewServer(constants.DefaultServerPort)
	err := server.StartServer()
	if err != nil {
		fmt.Println("Error creating server:", err)
	}
}
