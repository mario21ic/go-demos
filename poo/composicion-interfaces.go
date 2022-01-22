package main

import "fmt"

type Encrypter interface {
	Encrypt([]byte) string
}

type Decrypter interface {
	Decrypt(string, interface{}) error
}

type EncryptDecrypter interface {
	Encrypter
	Decrypter
}

func main() {
	g := gopher{animal{"Mario's Gopher"}, 4}
	g.Encrypt()
}

type animal struct {
	name string
}

type gopher struct {
	animal
	legs int
}

func (g gopher) Encrypt() {
	fmt.Printf("Encrypt: I'm %s and I'm a %T", g.name, g)
}
