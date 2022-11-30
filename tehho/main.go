package main

import (
	"fmt"
	"os"

	"rsc.io/quote"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	data, err := os.ReadFile("data/HelloWorld")
	check(err)
	fmt.Println(string(data))

	fmt.Println(HelloWorld())

	fmt.Println(quote.Go())
}
