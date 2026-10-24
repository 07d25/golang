package main

import "fmt"

func main() {
	age := 20
	if age >= 18 {
		fmt.Println("mayor de age")
	} else {
		fmt.Println("menor de age")
	}
	switch age {
	case 0:
		fmt.Println("recién nacido")
	case 18, 19, 20:
		fmt.Println("adulto joven")
	default:
		fmt.Println("otra age")
	}
}
