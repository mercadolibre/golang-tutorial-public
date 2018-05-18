// Package counter provides primitives for receiving a file and returning how
// many lines it contains.
package counter

import (
	"bufio"
	"os"
)

// CountFile receives a file and returns the number of lines in it.
func CountFile(f *os.File) int64 {
	var count int64

	s := bufio.NewScanner(f)
	for s.Scan() {
		count++
	}

	return count
}

// Reader is an interface compatible with io.Reader
type Reader interface {
	Read([]byte) (int, error)
}

// CountReader receives a reader and returns the number of lines in it.
func CountReader(r Reader) int64 {
	var count int64

	s := bufio.NewScanner(r)
	for s.Scan() {
		count++
	}

	return count
}
