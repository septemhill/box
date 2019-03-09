package box

import (
	"bytes"
	"fmt"
	"strings"

	es "github.com/septemhill/escapestring"
)

const DEFAULT_BORDER = "- │ ┌ ┐ └ ┘"
const esc = "\x1b["

func escape(str string) {
	fmt.Printf("%s%s", esc, str)
}

func moveToPaint(x, y int, str string) {
	escape(fmt.Sprintf("%d;%dH%s", x, y, str))
}

type Box struct {
	newbuff *bytes.Buffer
	oldbuff *bytes.Buffer
	X       int
	Y       int
	Width   int
	Height  int
	label   string
}

func (b *Box) Label() string {
	return b.label
}

func (b Box) Write(p []byte) (n int, err error) {
	return b.newbuff.Write(p)
}

func (b *Box) drawline(row, x, width int, es es.EscapeString) {
	for i := 0; i < width; i++ {
		moveToPaint(row, x+i, es.Element(i))
	}
}

func (b *Box) clearArea() {
	for i := b.Y; i < b.Height+b.Y; i++ {
		for j := b.X; j < b.Width+b.X; j++ {
			moveToPaint(i, j, " ")
		}
	}
}

func (b *Box) Draw() {
	var line string
	borders := strings.Split(DEFAULT_BORDER, " ")
	linecnt := 0

	if len(b.newbuff.Bytes()) == 0 {
		b.newbuff.Write(b.oldbuff.Bytes())
	} else {
		b.oldbuff.Reset()
		b.oldbuff.Write(b.newbuff.Bytes())
	}

	lines := bytes.Split(b.newbuff.Bytes(), []byte("\n"))

	b.clearArea()
	for i := b.Y; i < b.Height+b.Y; i++ {
		if i == b.Y {
			line += borders[2] + strings.Repeat(borders[0], b.Width-2) + borders[3]
		} else if i == b.Height+b.Y-1 {
			line += borders[4] + strings.Repeat(borders[0], b.Width-2) + borders[5]
		} else if linecnt < len(lines) {
			es := es.NewEscapeString(string(lines[linecnt]))
			if es.Width() > b.Width-2 {
				line += borders[1] + es.SubstringByWidth(0, b.Width-2) + borders[1]
			} else {
				line += borders[1] + es.String() + strings.Repeat(" ", b.Width-es.Width()-2) + borders[1]
			}
			linecnt++
		} else {
			line += borders[1] + strings.Repeat(" ", b.Width-2) + borders[1]
		}

		es := es.NewEscapeString(line)
		b.drawline(i, b.X, b.Width, es)

		line = ""
	}

	b.newbuff.Reset()
}

func NewBox(startX, startY, width, height int, label string) *Box {
	box := &Box{
		newbuff: bytes.NewBuffer(nil),
		oldbuff: bytes.NewBuffer(nil),
		X:       startX,
		Y:       startY,
		Width:   width,
		Height:  height,
		label:   label,
	}

	return box
}
