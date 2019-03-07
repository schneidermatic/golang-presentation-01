package main

import (
	"fmt"
)

type Author struct {
	Name string
}

type Book struct {
	*Author
	Title string
}

func (author *Author) Introduce() {
	fmt.Printf("Hi, I'm %s\n", author.Name)
}

func (book Book) Introduce() { // !!!
	fmt.Printf("Hi, I'm %s de Saint-Exupéry!!!\n", book.Name)
}

func main() {
	book := &Book{
		Author: Author{"Antoine"}, // !!!
		Title:  "The little Prince",
	}
	fmt.Printf("book.Name: %s\n", book.Name)
	fmt.Printf("book.Author.Name: %s\n", book.Author.Name)
	book.Introduce()
	book.Author.Introduce()
}
