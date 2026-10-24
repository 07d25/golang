package main

import "fmt"

func squareRoot(n float64) (float64, error) {
	if n < 0 {
		return 0, fmt.Errorf("%v no tiene raíz real", n)
	}
	return n, nil
}
func main() {
	_, err := squareRoot(-1)
	if err != nil {
		fmt.Println("error controlado:", err)
	}
}
