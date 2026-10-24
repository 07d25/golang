package main

import "fmt"

func dividir(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("no se puede dividir por cero")
	}
	return a / b, nil
}

func main() { resultado, err := dividir(10, 2); fmt.Println(resultado, err) }
