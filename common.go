package box

import (
	"os"

	"golang.org/x/sys/unix"
)

func getWindowSize() (int, int) {
	ws, err := unix.IoctlGetWinsize(int(os.Stdout.Fd()), unix.TIOCGWINSZ)

	if err != nil {
		return -1, -1
	}

	return int(ws.Col), int(ws.Row)
}
