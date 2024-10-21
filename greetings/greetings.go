package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Func: return hello-message (type string) + error
func Hello(name string) (string, error) {
	//If incorrect name ("empty name" error + empty name)
	if name == "" {
		return name, errors.New("empty name")
	}
	// Create a message (string type) + nil error
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func randomFormat() string {
	//Slice of "hello" messages
	formats := []string{
		"Hello, %v! Nice to see you!",
		"Great to see you, %v!",
		"Hi, %v!",
	}
	//Choose random one
	return formats[rand.Intn(len(formats))]
}
