module github.com/hijenhek/go_doc_tutorial/create_a_module/hello

go 1.18

replace github.com/hijenhek/go_doc_tutorial/create_a_module/greetings => ../greetings

require github.com/hijenhek/go_doc_tutorial/create_a_module/greetings v0.0.0-00010101000000-000000000000
