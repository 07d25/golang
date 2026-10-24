package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) Introduce() string {
	return fmt.Sprintf("Soy %s y tengo %d años", p.Name, p.Age)
}
func main() { fmt.Println((Person{Name: "Ana", Age: 30}).Introduce()) }
