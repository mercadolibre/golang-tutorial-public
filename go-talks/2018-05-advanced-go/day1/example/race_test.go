package race_test

import (
	"sync"
	"testing"
)

// START 1 OMIT
func TestRaceCounter(t *testing.T) {
	counter := new(int) // HLxxx
	var wg sync.WaitGroup
	// multiple producers
	for producerId := 0; producerId < 10; producerId++ {
		wg.Add(1)

		// spawn a new counter
		go func(producerId int) {
			defer wg.Done()
			*counter = *counter + 1 // HLxxx
		}(producerId)
	}

	// wait for the N producers
	wg.Wait()

	// check expected value
	if *counter != 10 {
		t.Fatal("counter failed.")
	}
}

// END 1 OMIT
