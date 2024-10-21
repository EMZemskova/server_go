package main

import (
	"fmt"
	"log"

	"github.com/EMZemskova/server_go/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0) //No flags output

	// A slice of names.
	names := []string{"Kate", "John", "Samanta"}

	// Get a greeting message and error
	messages, err := greetings.Hellos(names)

	//If error
	if err != nil {
		log.Fatal(err)
	}
	//Correct output
	fmt.Println(messages)
}
