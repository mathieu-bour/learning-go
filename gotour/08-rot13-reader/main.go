package main

import (
	"io"
	"os"
	"strings"
)

// Instructions: https://go.dev/tour/methods/23

type rot13Reader struct {
	r io.Reader
}

func (reader rot13Reader) Read(buffer []byte) (int, error) {
	var i = 0
	for {
		var data []byte = make([]byte, 8)
		l, err := reader.r.Read(data)

		if err == io.EOF || l == 0 {
			break
		}

		for _, letter := range data {
			if 97 <= letter && letter <= 122 {
				buffer[i] = ((letter-97)+13)%26 + 97
			} else if 65 <= letter && letter <= 90 {
				buffer[i] = ((letter-65)+13)%26 + 65
			} else {
				buffer[i] = letter
			}

			i++
		}
	}

	return i, nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
