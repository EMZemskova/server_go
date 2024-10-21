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

// Func: return hello-messages (type string) to many people + error
func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)
	for _, name := range names {
		message, err := Hello(name)

		if err != nil {
			return nil, err
		}
		messages[name] = message
	}

	return messages, nil
}

func randomFormat() string {
	//Slice of "hello" messages
	formats := []string{
		"Hello, %v! Nice to see you!",
		"Good to see you, %v!",
		"Hi, %v!",
	}
	//Choose random one
	return formats[rand.Intn(len(formats))]
}
