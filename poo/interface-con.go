package main

import "fmt"

type Article struct {
	Title  string
	Author string
}

func (a Article) String() string {
	return fmt.Sprintf("The %q article was written by %s.", a.Title, a.Author)
}

type Book struct {
	Title  string
	Author string
	Pages  int
}

func (b Book) String() string {
	return fmt.Sprintf("The %q book was written by %s. It has %d pages.", b.Title, b.Author, b.Pages)
}

func main() {
	a := Article{
		Title:  "Interfaces in Go",
		Author: "mario21ic",
	}

	Print(a)

	b := Book{
		Title:  "All About Go",
		Author: "Jenny Dolphin",
		Pages:  26,
	}
	Print(b)
}

// Definimos solo la funcion String()
type Stringer interface {
	String() string
}

// func Print(a Article) {
// 	fmt.Println(a.String())
// }
func Print(s Stringer) {
	fmt.Println(s.String())
}
