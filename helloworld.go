package main

import (
	"fmt"
	"os"
	"strings"
)

func GetCreatures () string {
	var dog = os.Getenv("DOG")
	var cat = os.Getenv("CAT")
	var synx = os.Getenv("SYNX")
	var creatures = []string{dog, cat, synx}
    return creatures[]
}
  

func main() {
	GetCreatures(creatures)
	fmt.Println("Hello ", strings.Join(creatures, ", "), "!")
}
