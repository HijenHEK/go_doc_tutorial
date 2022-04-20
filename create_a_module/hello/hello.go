package main

import (
	"fmt"
	"log"

	"github.com/hijenhek/go_doc_tutorial/create_a_module/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	message, err := greetings.Hello("")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}
