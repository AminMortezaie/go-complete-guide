package main

import "fmt"

type Book struct {
	ID        int
	title     string
	shortDesc string
	price     int
}

func main() {
	firstBook := *NewBook(1, "num1", "num1 is good", 100)
	firstBook.printData()
}

func NewBook(ID int, title string, shortDesc string, price int) (newBook *Book) {
	newBook = &Book{ID, title, shortDesc, price}
	// fmt.Println("this might be reference", newBook)
	return
}

func (book *Book) printData() {
	fmt.Printf("%v %v %v", book.title, book.shortDesc, book.price)

}
