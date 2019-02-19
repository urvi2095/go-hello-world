package functions

import "errors"

const HelloWorldMessage = "Hello world public constant"

func GetMessage() (string, string) {
	message := "hi"
	name := "urvi"
	return message, name
}

func CreateMessage(x int) (int, error) {

	if x < 0 {
		return  0, errors.New("value of x cannot be less than 0")
	}
	number := x
	return number, nil
}
