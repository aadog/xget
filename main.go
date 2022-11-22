package main

import (
	"crypto/md5"
	"fmt"
)

func main() {
	fmt.Println("hello2 3")
	m := md5.New()
	fmt.Print(m.Size())
}
