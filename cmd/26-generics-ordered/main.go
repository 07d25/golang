package main

import "fmt"

type Ordered interface {
	~int | ~int64 | ~float64 | ~string
}

func Mayor[T Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}
func main() { fmt.Println(Mayor(4, 9), Mayor("go", "rust")) }
