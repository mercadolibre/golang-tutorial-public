package channels_test

import "testing"

func TestBlockingChannel(t *testing.T) {

	// START 1 OMIT
	select {
	case ch <- 42: // HLxxx
		// non-blocking write
	default:
		// sin default habría bloqueao
	}
	// END 1 OMIT

	// START 2 OMIT
	select {
	case value := <-ch: // HLxxx
		// non-blocking read
	default:
		// sin default habría bloqueao
	}
	// END 2 OMIT
}
