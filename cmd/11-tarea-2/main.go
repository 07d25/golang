package main

import "fmt"

func main() {
	numbers := []int{2, 7, 4, 9, 6}
	for _, n := range numbers {
		if n%2 == 0 {
			fmt.Println(n, "es par")
		}
	}
}
