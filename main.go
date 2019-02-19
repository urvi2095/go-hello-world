package main

import (
	"fmt"
	"github.com/go-hello-world/functions"
	"log"
)

func main() {

	var msg string = "Hello World!!"
	println(msg)

	m, _ := functions.GetMessage()
	if v, error := functions.CreateMessage(-2); error != nil {
		fmt.Printf(" error: %v\n", error)
	}   else {
		fmt.Printf("%v\n", v)
	}
	log.Printf("%v", m)
}

//func getMessage() (string, string) {
//	message := "hi"
//	name := "urvi"
//	return message, name
//}
