package main

import "fmt"

func Contiene[T comparable](elementos []T, buscado T) bool {
	for _, elemento := range elementos {
		if elemento == buscado {
			return true
		}
	}
	return false
}
func main() { fmt.Println(Contiene([]int{1, 2, 3}, 2), Contiene([]string{"go"}, "rust")) }
