package main

import "fmt"

func main() {
	sales := map[string]float64{"enero": 120, "febrero": 180, "marzo": 150}
	total := 0.0
	for month, sale := range sales {
		fmt.Println(month, sale)
		total += sale
	}
	fmt.Println("total:", total)
}
