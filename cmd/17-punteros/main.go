package main

import "fmt"

func doubleValue(valor *int) { *valor *= 2 }
func main()                  { number := 7; doubleValue(&number); fmt.Println(number) }
