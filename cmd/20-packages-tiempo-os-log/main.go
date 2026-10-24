package main

import (
	"log"
	"os"
	"time"
)

func main() { log.SetOutput(os.Stdout); log.Println("hora local:", time.Now().Format(time.RFC3339)) }
