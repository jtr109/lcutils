package acm

import (
	"io/ioutil"
	"os"
)

func MockStdout() (pipeReader, pipeWriter, oldStdout *os.File, err error) {
	oldStdout = os.Stdout
	pipeReader, pipeWriter, err = os.Pipe()
	if err != nil {
		return nil, nil, nil, err
	}
	os.Stdout = pipeWriter
	return
}

func CaptureStdout(pipeReader, pipeWriter, oldStdout *os.File) (string, error) {
	defer func() { os.Stdout = oldStdout }() // restore original Stdout
	pipeWriter.Close()
	out, _ := ioutil.ReadAll(pipeReader)
	return string(out), nil
}
