package main

import (
	"fmt"
)

func main() {
	a := "aaa"

	b := &a

	name(b)

}

func name(n *int) {

	fmt.Println(n)
}
