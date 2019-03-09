package box

import "errors"

type BoxManager struct {
	boxes  map[string]*Box
	curBox *Box
}

func (bm *BoxManager) AddBox(b *Box) error {
	if _, ok := bm.boxes[b.Label()]; ok {
		return errors.New("box alread existed")
	}

	bm.boxes[b.Label()] = b

	return nil
}

func (bm *BoxManager) FocusBox(label string) error {
	if _, ok := bm.boxes[label]; !ok {
		return errors.New("non-exist box")
	}

	bm.boxes[label].Draw()

	return nil
}

func NewBoxManager() *BoxManager {
	bm := &BoxManager{
		boxes:  make(map[string]*Box),
		curBox: nil,
	}

	return bm
}
