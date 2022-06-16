package main

import "fmt"

type Book struct {
	title   string
	author  string
	subject string
	book_id int
}

func main() {

	var book Book

	book.author = "PLO LUMUMBA"
	book.book_id = 22
	book.subject = "Chemistry"
	book.title = "KLB"

	printBook(&book)

}

func printBook(book *Book) {
	fmt.Printf("Book title is: %s\n", book.title)
	fmt.Printf("Book author is: %s\n", book.author)
	fmt.Printf("Book subject is: %s\n", book.subject)
	fmt.Printf("Book id is: %d\n", book.book_id)

}
