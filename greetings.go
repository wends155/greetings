package greetings

import "fmt"

// add comment
func Hello(name string) string {
	message := fmt.Sprintf("Hi, %v. Welcome Aboard !", name)
	return message
}
