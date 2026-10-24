package main

import "fmt"

func main() {
	a, b := uint8(0b1100), uint8(0b1010)
	fmt.Printf("AND=%04b OR=%04b XOR=%04b NOT=%08b\n", a&b, a|b, a^b, ^a)
}
