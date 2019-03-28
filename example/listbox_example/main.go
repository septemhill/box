package main

import (
	"fmt"
	"time"

	"github.com/septemhill/box"
)

func clearScreen() {
	fmt.Println("\x1b[2J")
}

func main() {
	clearScreen()
	l1 := box.NewListBox(1, 1, 20, 15, "l1", []string{
		"Amazon",
		"Apple",
		"Facebook",
		"Google",
		"Microsoft",
	})
	l1.Draw()

	for i := 0; i < 5; i++ {
		l1.ArrowControl(box.DOWN_ARROW)
		l1.Draw()
		time.Sleep(time.Millisecond * 300)
	}

	for i := 0; i < 5; i++ {
		l1.ArrowControl(box.UP_ARROW)
		l1.Draw()
		time.Sleep(time.Millisecond * 300)
	}

	fmt.Println()
}
