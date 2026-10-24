package main

import "fmt"

func square(n int, output chan<- int) { output <- n * n }
func main()                           { output := make(chan int); go square(5, output); fmt.Println(<-output); close(output) }
