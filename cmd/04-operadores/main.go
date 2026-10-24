package main

import "fmt"

func main() {
	a, b := 10, 3
	fmt.Println("aritméticos:", a+b, a-b, a*b, a/b, a%b)
	fmt.Println("comparación:", a > b, a == b)
	fmt.Println("lógicos:", a > 0 && b > 0, a == 0 || b == 0, !(a == b))
}
