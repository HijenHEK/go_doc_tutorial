package main

import (
	"fmt"
	"log"

	"github.com/hijenhek/go_doc_tutorial/create_a_module/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{"Hijen", "Hajed", "Hiba"}
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(messages)

}
