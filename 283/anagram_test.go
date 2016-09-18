package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"testing"
)

func TestAnagram(t *testing.T) {
	in, err := os.Open("input.txt")
	if err != nil {
		t.Error(err)
	}

	op, err := ioutil.ReadFile("output.txt")
	if err != nil {
		t.Error(err)
	}
	benchmark := string(op)

	out := captureOutput(func() {
		runner(in)
	})
	if benchmark != out {
		t.Errorf("expected:\n%s\ngot:\n%s", benchmark, out)
	}
}

func captureOutput(f func()) string {
	var buf bytes.Buffer

	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	os.Stdout = old
	io.Copy(&buf, r)
	return buf.String()
}
