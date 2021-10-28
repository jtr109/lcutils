package acm

import (
	"os"
)

func MockStdin(input string) (oldStdin *os.File, err error) {
	content := []byte(input)
	r, w, err := os.Pipe()
	if err != nil {
		return nil, err
	}

	if _, err = w.Write(content); err != nil {
		return nil, err
	}
	w.Close()

	oldStdin = os.Stdin
	os.Stdin = r
	return
}

func RestoreStdin(oldStdin *os.File) {
	os.Stdin = oldStdin // restore operate system standard input
}
