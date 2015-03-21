package main

import (
	"io"
	"os"
)

func main() {
	// Open terminal
	fd, openerr := os.Open("/dev/tty0")
	if openerr != nil {
		panic(openerr)
	}
	// Close fd when done
	defer func() {
		if err := fd.Close(); err != nil {
			panic(err)
		}
	}()
	bel := []byte{7}
	_, writeerr := fd.Write(bel)
}
