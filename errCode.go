package box

import "errors"

var (
	ERR_BOX_OUT_OF_RANGE  = errors.New("cannot out of window")
	ERR_BOX_NON_EXIST     = errors.New("non-exist box")
	ERR_BOX_ALREADY_EXIST = errors.New("box already existed")
	ERR_BOX_WINDOW_SIZE   = errors.New("cannot get window size")
)
