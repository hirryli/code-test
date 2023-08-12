package main

import (
	"fmt"
	"testgo/test"
)

type BB interface {
	SetName(string)
	GetName() string
	SetTotal(int)
	GetTotal() int
}

func main() {
	fmt.Println("Hello the world!")
	var book test.Book
	book.Name = "test"
	var tt BB
	tt = book
	// var tt test.Book
	// tt.SetName("tetsst")
	fmt.Println("tt.Name:" + tt.GetName())
}
