package main

import "fmt"

func main() {
	namonth := []string{"Ana", "Luis"}
	namonth = append(namonth, "Marta")
	parte := namonth[1:]
	fmt.Println(namonth, parte, len(namonth), cap(namonth))
}
