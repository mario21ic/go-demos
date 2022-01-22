package main

import "fmt"

type Article struct {
	Title  string
	Author string
}

func (a Article) String() string {
	return fmt.Sprintf("The %q article was written by %s.", a.Title, a.Author)
}

func main() {
	a := Article{
		Title:  "Interfaces in Go",
		Author: "mario21ic",
	}
	fmt.Println(a.String())
}
