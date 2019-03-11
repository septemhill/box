package box

type BoxManager struct {
	boxes  map[string]*Box
	curBox *Box
	width  int
	height int
}

func (bm *BoxManager) focus(label string) {
	bm.curBox = bm.boxes[label]
	bm.boxes[label].Draw()
}

func (bm *BoxManager) AddBox(b *Box) error {
	if _, ok := bm.boxes[b.Label()]; ok {
		return ERR_BOX_ALREADY_EXIST
	}

	bm.boxes[b.Label()] = b
	bm.focus(b.Label())

	return nil
}

func (bm *BoxManager) FocusBox(label string) error {
	if _, ok := bm.boxes[label]; !ok {
		return ERR_BOX_NON_EXIST
	}

	bm.focus(label)

	return nil
}

func (bm *BoxManager) MoveBox(label string, x, y int) error {
	if _, ok := bm.boxes[label]; !ok {
		return ERR_BOX_NON_EXIST
	}

	b := bm.boxes[label]

	if (x+b.Width < bm.width) && (y+b.Height < bm.height) {
		b.X = x
		b.Y = y
		bm.focus(label)
	}

	return ERR_BOX_OUT_OF_RANGE
}

func (bm *BoxManager) ResizeBox(label string, width, height int) error {
	if _, ok := bm.boxes[label]; !ok {
		return ERR_BOX_NON_EXIST
	}

	b := bm.boxes[label]

	if (width+b.X < bm.width) && (height+b.Y < bm.height) {
		b.Width = width
		b.Height = height
		bm.focus(label)
	}

	return ERR_BOX_OUT_OF_RANGE
}

func NewBoxManager() (*BoxManager, error) {
	bm := &BoxManager{
		boxes:  make(map[string]*Box),
		curBox: nil,
	}

	bm.width, bm.height = getWindowSize()

	if bm.width < 0 || bm.height < 0 {
		return nil, ERR_BOX_WINDOW_SIZE
	}

	return bm, nil
}
