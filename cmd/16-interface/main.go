package main

import "fmt"

type Sound interface{ MakeSound() string }
type Dog struct{}

func (Dog) MakeSound() string { return "guau" }
func playSound(s Sound)       { fmt.Println(s.MakeSound()) }
func main()                   { playSound(Dog{}) }
