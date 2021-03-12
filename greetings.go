package greetings

import "fmt"

// add comment .. testing changes ... testing for ssh push
func Hello(name string) string {
	message := fmt.Sprintf("Hi, %v. Welcome Aboard !", name)
	return message
}
