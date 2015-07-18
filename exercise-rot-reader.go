package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (this *rot13Reader) Read(bytes []byte) (int, error) {
	n, err := this.r.Read(bytes)
	for i := range bytes {
		c := bytes[i]
		switch {
		case c >= 'a' && c <= 'm' || c >= 'A' && c <= 'M': c += 13
		case c >= 'n' && c <= 'z' || c >= 'N' && c <= 'Z': c -= 13
		}
		bytes[i] = c
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
