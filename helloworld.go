package main

import (
	"fmt"
	"os"
	"strings"
)

func GetCreatures() []string {
	var dog = os.Getenv("DOG")
	var cat = os.Getenv("CAT")
	var synx = os.Getenv("SYNX")
	var beings = []string{dog, cat, synx}
	return beings
}

func main() {
	var creatures = GetCreatures()
	fmt.Println("Hello", strings.Join(creatures, ", "), "!")
	fmt.Println(strings.Join(creatures[0:2], ", "), "and", creatures[2], "are my favorite")
}
