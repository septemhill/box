package box

import (
	"bytes"
	"fmt"
	"strings"

	es "github.com/septemhill/escapestring"
	"github.com/septemhill/fion"
)

const (
	UP_ARROW = iota
	DOWN_ARROW
	LEFT_ARROW
	RIGHT_ARROW
)

type ListBox struct {
	*Box
	list   []string
	curpos int
}

func (l *ListBox) ArrowControl(ctrlCode int) {
	switch ctrlCode {
	case UP_ARROW:
		if l.curpos-1 < 0 {
			l.curpos = len(l.list) - 1
		} else {
			l.curpos--
		}
	case DOWN_ARROW:
		if l.curpos+1 >= len(l.list) {
			l.curpos = 0
		} else {
			l.curpos++
		}
	default:
		return
	}

	l.Draw()
}

func (l *ListBox) Draw() {
	var line string
	borders := strings.Split(DEFAULT_BORDER, " ")
	linecnt := 0

	if len(l.newbuff.Bytes()) == 0 {
		l.newbuff.Write(l.oldbuff.Bytes())
	} else {
		l.oldbuff.Reset()
		l.oldbuff.Write(l.newbuff.Bytes())
	}

	lines := bytes.Split(l.newbuff.Bytes(), []byte("\n"))

	//l.clearArea()
	for i := l.Y; i < l.Height+l.Y; i++ {
		if i == l.Y {
			line += borders[2] + strings.Repeat(borders[0], l.Width-2) + borders[3]
		} else if i == l.Height+l.Y-1 {
			line += borders[4] + strings.Repeat(borders[0], l.Width-2) + borders[5]
		} else if linecnt < len(lines) {
			es := es.NewEscapeString(string(lines[linecnt]))
			if es.Width() > l.Width-2 {
				if l.curpos == linecnt {
					line += borders[1] + fion.BRed(es.SubstringByWidth(0, l.Width-2)) + borders[1]
				} else {
					line += borders[1] + es.SubstringByWidth(0, l.Width-2) + borders[1]
				}
			} else {
				if l.curpos == linecnt {
					line += borders[1] + fion.BRed(es.String()+strings.Repeat(" ", l.Width-es.Width()-2)) + borders[1]
				} else {
					line += borders[1] + es.String() + strings.Repeat(" ", l.Width-es.Width()-2) + borders[1]
				}
			}
			linecnt++
		} else {
			line += borders[1] + strings.Repeat(" ", l.Width-2) + borders[1]
		}

		es := es.NewEscapeString(line)
		l.drawline(i, l.X, l.Width, es)

		line = ""
	}

	l.newbuff.Reset()
}

func (l *ListBox) updateList() {
	var rstr string

	for i := 0; i < len(l.list); i++ {
		rstr += l.list[i] + "\n"
	}

	fmt.Fprintf(l.Box, rstr)
	//l.Draw()
}

func NewListBox(startX, startY, width, height int, label string, list []string) *ListBox {
	l := &ListBox{
		Box: NewBox(startX, startY, width, height, label),
	}

	l.list = make([]string, len(list))
	copy(l.list, list)
	l.curpos = 0

	l.updateList()

	return l
}
