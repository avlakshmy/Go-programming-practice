/*Implementation of a rot13Reader: implements io.Reader and reads from an io.Reader, modifying the stream by applying the rot13 substitution cipher to all alphabetical characters.*/

package main

import (
	"io"
	"os"
	"strings"
	"unicode"
)

type rot13Reader struct {
	r io.Reader
}

func (rr rot13Reader) Read (b []byte) (int, error) {
	n, err := rr.r.Read(b)
	for i := 0; i < n; i++ {
		c := rune(b[i])
		if unicode.IsUpper(c) {
			k := b[i] - 65 
			if k < 13 {
				b[i] = b[i] + 13
			} else {
				b[i] = 65 + (k+13)%26 
			}
		} else {
			if unicode.IsLower(c) {
				k := b[i] - 97 
				if k < 13 {
					b[i] = b[i] + 13
				} else {
					b[i] = 97 + (k+13)%26 
				}
			} 	
		}
	}
	return n, err
}

func main() {
	s := strings.NewReader("Yrneavat Tb")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

