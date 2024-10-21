package greetings

import "fmt"

//Func: return hello-message (type string)
func Hello(name string) string {
	// Create a message (string type)
	message := fmt.Sprintf("Hi, %v. Welcome to the server!", name)
	return message
}
