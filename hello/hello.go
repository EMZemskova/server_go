package main

import (
	"fmt"
	"log"

	"github.com/EMZemskova/server_go/greetings"
)

func main() {

	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// Get a greeting message and error
	message, err := greetings.Hello("")
	//If error
	if err != nil {
		log.Fatal(err)
	}
	//Correct output
	fmt.Println(message)
}
