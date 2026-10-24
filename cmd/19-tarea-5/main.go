package main

import "fmt"

func average(grades ...float64) (float64, error) {
	if len(grades) == 0 {
		return 0, fmt.Errorf("se requiere al menos una grade")
	}
	total := 0.0
	for _, n := range grades {
		total += n
	}
	return total / float64(len(grades)), nil
}
func main() { average, err := average(8, 9.5, 10); fmt.Println(average, err) }
