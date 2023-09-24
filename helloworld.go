package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func GetCreatures() []string {
	var dog = os.Getenv("DOG")
	var cat = os.Getenv("CAT")
	var synx = os.Getenv("SYNX")
	var beings = []string{dog, cat, synx}
	return beings
}

func main() {
	var time = time.Now()
	var creatures = GetCreatures()
	fmt.Println("Hello", strings.Join(creatures, ", "), "!")
	fmt.Println(strings.Join(creatures[0:2], ", "), "and", creatures[2], "are my favorite")
	fmt.Println(creatures[2], "and", creatures[1], "are my partners in crime!")
	fmt.Println(creatures[0], "is a very good pawgrammer doggie")
	fmt.Println("the current time is", time.Format("01-02-2006 15:04:05"))
}
