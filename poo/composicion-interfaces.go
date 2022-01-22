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
	// g.Encrypt()
	// g.Decrypt()
	Print(g)
}

func Print(m EncryptDecrypter) {
	// fmt.Println(m.Encrypt())
	fmt.Println(m.Decrypt("xD"))
}

type animal struct {
	name string
}

type gopher struct {
	animal
	legs int
}

func (g gopher) Encrypt(b []byte) {
	fmt.Printf("Encrypt: I'm %s and I'm a %T\n", g.name, g)
}

func (g gopher) Decrypt(s string) {
	fmt.Printf("Decrypt: I'm %s and I'm a %T\n", g.name, g)
}
