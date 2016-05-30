package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func rot13(x byte) byte {
	captial := (x >= 'A' && x <= 'Z')
	if !captial && (x < 'a' || x > 'z') {
		return x
	}
	x += 13
	if captial && x > 'Z' || !captial && x > 'z' {
		return x - 26
	}

	return x
}

func (rot13r *rot13Reader) Read(b []byte) (int, error) {
	bInternal := make([]byte, len(b))
	n, err := rot13r.r.Read(bInternal)
	if err == io.EOF {
		return 0, io.EOF
	}
	for i := range bInternal {
		b[i] = rot13(bInternal[i])
	}
	return n, nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!\n")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
