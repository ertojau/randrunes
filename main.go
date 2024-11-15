//input: ask to user for a number and a string
//use the cadena package
//output: print the output of package cadena

package main

import (
	"fmt"
	"github.com/ertojau/randrunes/cadena"
	"os"
)

func main() {

	var num int
	var str string
	fmt.Print("Enter a number: ")
	fmt.Scan(&num)
	fmt.Print("Enter a string: ")
	fmt.Scan(&str)
	fmt.Println("Your string is: ", cadena.Cadena(str, num))
}