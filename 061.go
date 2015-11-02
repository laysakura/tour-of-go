package main
import (
	"io"
	"strings"
	"os"
)

type rot13Reader struct  {
	r io.Reader
}

func (rot *rot13Reader) Read(p []byte) (int, error) {
	n, _ := rot.r.Read(p)

	for i := 0; i < n; i++ {
		if ( ('A' <= p[i] && p[i] <= 'M') || ('a' <= p[i] && p[i] <= 'm') ) {
			p[i] += 13
		} else if ( ('N' <= p[i] && p[i] <= 'Z') || ('n' <= p[i] && p[i] <= 'z') ) {
			p[i] -= 13
		}
	}
	return n, nil
}

func main()  {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
