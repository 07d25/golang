package main

import "fmt"

func main() {
	var letter byte = 'G' // byte es un alias de uint8; representa un byte, no siempre un carácter completo.
	monthsage := []byte("Go")
	fmt.Println(letter, string(monthsage), monthsage)
}
