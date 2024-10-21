package greetings

import (
	"errors"
	"fmt"
)

// Func: return hello-message (type string) + error
func Hello(name string) (string, error) {
	//If incorrect name ("empty name" error + empty name)
	if name == "" {
		return " ", errors.New("empty name")
	}
	// Create a message (string type) + nil error
	message := fmt.Sprintf("Hi, %v. Welcome to the server!", name)
	return message, nil
}
