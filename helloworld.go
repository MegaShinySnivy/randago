package main

import (
	"fmt"
	"strings"
)

func main() {
	var dog = "Ebu"
	var cat = "Kat"
	var synx = "Jinx"
	var creatures = []string{dog, cat, synx}
	fmt.Println("Hello ", strings.Join(creatures, ", "), "!")
}
