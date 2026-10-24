package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// Declaración explícita de variables con distintos tipos.
	var integerValue int = 5
	fmt.Println(integerValue)
	var unsignedValue uint = 15
	fmt.Println(unsignedValue)
	var stringValue string = "My string variable"
	fmt.Println(stringValue)
	var boolValue bool = true
	fmt.Println(boolValue)

	// Con := Go infiere el tipo dentro de una función.
	name := "Martin"
	fmt.Println(name)
	fmt.Println(false)
	fmt.Println(25)

	// Las constantes no pueden modificarse después de declararse.
	const countryName = "Argentina"
	fmt.Println(countryName)

	var eightBitValue int8
	fmt.Printf("Valor cero de int8: %d\n", eightBitValue)
	fmt.Printf("tipo: %T, bytes: %d, bits: %d\n", eightBitValue, unsafe.Sizeof(eightBitValue), unsafe.Sizeof(eightBitValue)*8)
	var sixteenBitValue int16
	fmt.Printf("Valor cero de int16: %d\n", sixteenBitValue)
	fmt.Printf("tipo: %T, bytes: %d\n", sixteenBitValue, unsafe.Sizeof(sixteenBitValue))
	var emptyString string
	fmt.Printf("Valor cero de string: %q\n", emptyString)
}
