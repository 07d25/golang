package main

import "fmt"

func main() {
	var grades [3]int = [3]int{8, 9, 10}
	for index, grade := range grades {
		fmt.Printf("grade[%d]=%d\n", index, grade)
	}
}
