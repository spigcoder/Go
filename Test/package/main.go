package main

import(
	"Test/add"
	"fmt"
)

func main(){
	a, b := 1, 2
	c := add.Add(a, b)
	if c < add.MAX{
		fmt.Println(c)
	}
}