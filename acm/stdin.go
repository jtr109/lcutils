package acm

import (
	"io/ioutil"
	"os"
)

func MockStdin(input string) (mockStdin, oldStdin *os.File, err error) {
	content := []byte(input)
	mockStdin, err = ioutil.TempFile("", "example")
	if err != nil {
		return nil, nil, err
	}

	if _, err := mockStdin.Write(content); err != nil {
		return mockStdin, nil, err
	}

	if _, err := mockStdin.Seek(0, 0); err != nil {
		return mockStdin, nil, err
	}

	oldStdin = os.Stdin
	os.Stdin = mockStdin
	return
}

func RestoreStdin(mockStdin, oldStdin *os.File) {
	os.Remove(mockStdin.Name()) // clean up tempfile
	os.Stdin = oldStdin         // restore system standard input
}
