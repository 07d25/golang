package main

import "fmt"

func main() {
	prices := map[string]float64{"pan": 1.5, "café": 2.25}
	prices["té"] = 1.75
	price, exists := prices["café"]
	fmt.Println(price, exists, prices)
}
