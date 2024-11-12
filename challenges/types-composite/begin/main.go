// challenges/types-composite/begin/main.go
package main

import (
	"fmt"
)

// define an author type with a name
type author struct {
	name string
}

// define a book type with a title and author
type book struct {
	title  string
	author string
}

// define a library type to track a list of books
type library struct {
	books []book
}

// define addBook to add a book to the library
func (l *library) addBook(title string, author string) string {
	newBook := book{title: title, author: author}
	l.books = append(l.books, newBook)
	return title + " was added to the library"
}

// define a lookupByAuthor function to find books by an author's name
func (l library) lookupByAuthor(author string) []book {
	var result []book
	for _, book := range l.books {
		if book.author == author {
			result = append(result, book)
		}
	}
	fmt.Println(result)
	return result
}

// func (l library) lookupByAuthor(author string)

func main() {
	// create a new library
	library1 := library{}
	// fmt.Print(library1)

	// add 2 books to the library by the same author
	library1.books = append(library1.books,
		book{title: "The Discourses", author: "Marcus"},
		book{title: "Letters from a Stoic", author: "Marcus"})

	// fmt.Print(library1)

	// add 1 book to the library by a different author
	library1.books = append(library1.books, book{title: "Imagine", author: "Telius"})

	// dump the library with spew
	// spew.Dump(library1)

	// lookup books by known author in the library
	library1.lookupByAuthor("Marcus")
	// print out the first book's title and its author's name
	fmt.Printf("The first book in the library is %s by %s\n", library1.books[0].title, library1.books[0].author)
}
