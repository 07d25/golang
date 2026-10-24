package main

import "fmt"

func main() {
	integer := 42
	decimalValue := float64(integer) // Conversión explícita entre tipos numéricos.
	textValue := fmt.Sprintf("%d", integer)
	fmt.Printf("integer=%d decimalValue=%.1f textValue=%q\n", integer, decimalValue, textValue)
}
