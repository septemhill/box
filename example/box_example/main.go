package main

import (
	"fmt"

	"github.com/septemhill/box"
)

func clearScreen() {
	fmt.Println("\x1b[2J")
}

func main() {
	b1 := box.NewBox(1, 1, 20, 10, "box1")
	fmt.Fprintf(b1, "Hi, Septem")
	b1.Draw()
}
