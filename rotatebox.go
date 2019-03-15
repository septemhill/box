package box

import (
	"bytes"
	"fmt"
	"strings"

	es "github.com/septemhill/escapestring"
	"github.com/septemhill/fion"
)

type RotateBox struct {
	*ListBox
	listCount int
}

func (r *RotateBox) ArrowControl(ctrlCode int) {
	switch ctrlCode {
	case UP_ARROW:
		if r.curpos-1 < 0 {
			r.curpos = len(r.list) - 1
		} else {
			r.curpos--
		}
	case DOWN_ARROW:
		if r.curpos+1 >= len(r.list) {
			r.curpos = 0
		} else {
			r.curpos++
		}
	default:
		return
	}

	r.updateList()
	r.Draw()
}

func (r *RotateBox) Draw() {
	var line string
	borders := strings.Split(DEFAULT_BORDER, " ")
	linecnt := 0

	if len(r.newbuff.Bytes()) == 0 {
		r.newbuff.Write(r.oldbuff.Bytes())
	} else {
		r.oldbuff.Reset()
		r.oldbuff.Write(r.newbuff.Bytes())
	}

	lines := bytes.Split(r.newbuff.Bytes(), []byte("\n"))

	for i := r.Y; i < r.Height+r.Y; i++ {
		if i == r.Y {
			line += borders[2] + strings.Repeat(borders[0], r.Width-2) + borders[3]
		} else if i == r.Height+r.Y-1 {
			line += borders[4] + strings.Repeat(borders[0], r.Width-2) + borders[5]
		} else if linecnt < len(lines) {
			es := es.NewEscapeString(string(lines[linecnt]))
			if es.Width() > r.Width-2 {
				if linecnt == len(lines)/2 {
					line += borders[1] + fion.BRed(es.String()+strings.Repeat(" ", r.Width-es.Width()-2)) + borders[1]
				} else {
					line += borders[1] + es.String() + strings.Repeat(" ", r.Width-es.Width()-2) + borders[1]
				}
			} else {
				if linecnt == len(lines)/2 {
					line += borders[1] + fion.BRed(es.String()+strings.Repeat(" ", r.Width-es.Width()-2)) + borders[1]
				} else {
					line += borders[1] + es.String() + strings.Repeat(" ", r.Width-es.Width()-2) + borders[1]
				}
			}
			linecnt++
		} else {
			line += borders[1] + strings.Repeat(" ", r.Width-2) + borders[1]
		}

		es := es.NewEscapeString(line)
		r.drawline(i, r.X, r.Width, es)

		line = ""
	}

	r.newbuff.Reset()
}

func (r *RotateBox) rotate() []string {
	list := make([]string, 0)
	nIdx := (r.curpos - r.listCount/2 + len(r.list)) % len(r.list)

	for i := 0; i < r.listCount; i++ {
		list = append(list, r.list[(nIdx+i)%len(r.list)])
	}

	return list
}

func (r *RotateBox) updateList() {
	var rstr string

	list := r.rotate()

	for i := 0; i < len(list); i++ {
		if i != len(list)-1 {
			rstr += list[i] + "\n"
		} else {
			rstr += list[i]
		}
	}

	r.newbuff.Reset()
	fmt.Fprintf(r.ListBox, rstr)
}

func NewRotateBox(startX, startY, width, listCount int, label string, list []string) (*RotateBox, error) {
	if listCount < 0 || listCount%2 == 0 || listCount > len(list) {
		return nil, ERR_BOX_LIST_COUNT
	}

	r := &RotateBox{
		ListBox:   NewListBox(startX, startY, width, listCount+2, label, list),
		listCount: listCount,
	}

	r.updateList()

	return r, nil
}
