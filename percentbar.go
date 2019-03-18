package box

import (
	"fmt"
	"math"
	"strings"
)

const (
	defaultBar = " ▏ ▎ ▍ ▌ ▋ ▊ ▉ █"
	barStage   = 8
	totalBars  = 12
)

type PercentBar struct {
	maxValue int
	curValue int
}

func (p *PercentBar) Draw() {
	var rstr string
	bars := strings.Split(defaultBar, " ")
	unit := float64(100) / (float64(barStage) * float64(totalBars))
	fullunit := unit * 8
	curpercent := float64(p.curValue) / float64(p.maxValue) * 100

	unitcnt := curpercent / unit
	fullcnt := int(unitcnt / fullunit)
	rmunitcnt := unitcnt - float64(fullcnt)*fullunit

	rstr += strings.Repeat(bars[8], fullcnt)
	rstr += bars[int(math.Floor(rmunitcnt/unit+0.5))]

	fmt.Println(rstr)
}

func (p *PercentBar) Add(i uint) {
	if p.curValue+int(i) > p.maxValue {
		p.curValue = p.maxValue
	} else {
		p.curValue += int(i)
	}

	p.Draw()
}

func (p *PercentBar) Sub(i uint) {
	if p.curValue-int(i) < 0 {
		p.curValue = 0
	} else {
		p.curValue -= int(i)
	}

	p.Draw()
}

func NewPercentBar(max int) *PercentBar {
	b := &PercentBar{
		maxValue: max,
		curValue: max,
	}

	return b
}
