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
	r1, _ := box.NewRotateBox(1, 1, 20, 5, "l1", []string{
		"AMAZ",
		"APPL",
		"FB",
		"GOOG",
		"MSFT",
		"VTI",
		"VWO",
		"BND",
		"VNQ",
		"VPL",
		"VGK",
	})
	r1.Draw()

	for i := 0; i < 5; i++ {
		r1.ArrowControl(box.DOWN_ARROW)
		r1.Draw()
		time.Sleep(time.Millisecond * 300)
	}

	time.Sleep(time.Second)

	for i := 0; i < 5; i++ {
		r1.ArrowControl(box.UP_ARROW)
		r1.Draw()
		time.Sleep(time.Millisecond * 300)
	}

	fmt.Println()
}
