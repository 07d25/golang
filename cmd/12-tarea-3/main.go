package main

import "fmt"

func main() {
	words := []string{"go", "es", "simple"}
	for i := range words {
		words[i] = words[i] + "!"
	}
	fmt.Println(words)
}
