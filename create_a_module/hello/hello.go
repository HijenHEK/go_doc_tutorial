package main

import (
	"fmt"

	"github.com/hijenhek/go_doc_tutorial/create_a_module/greetings"
)

func main() {
	message := greetings.Hello("Hijen")
	fmt.Println(message)
}
