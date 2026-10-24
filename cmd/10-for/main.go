package main

import "fmt"

func main() {
	sum := 0
	for i := 1; i <= 5; i++ {
		sum += i
	}
	for sum > 12 {
		sum--
		if sum == 14 {
			break
		}
	}
	fmt.Println("sum:", sum)
}
