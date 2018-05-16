package sync_primitives_test

import (
	"golang.org/x/sync/errgroup"
	"sync"
	"sync/atomic"
	"testing"
)

// START 1 OMIT
func TestRaceCounterFixed(t *testing.T) {
	counter := new(int)
	var mu sync.Mutex // HLxxx
	var wg sync.WaitGroup

	for producerId := 0; producerId < 10; producerId++ {
		wg.Add(1)
		// spawn a new counter
		go func(producerId int) {
			mu.Lock()         // HLxxx
			defer mu.Unlock() // HLxxx
			defer wg.Done()
			*counter = *counter + 1
		}(producerId)
	}
	wg.Wait()           // wait for the N producers
	if *counter != 10 { // check expected value
		t.Fatal("counter failed.")
	}
}

// END 1 OMIT

func TestMutexPrimitive1(t *testing.T) {
	// START 2 OMIT
	type Stats struct {
		Reads  int64
		Writes int64
	}
	// END 2 OMIT
}
func TestMutexPrimitive2(t *testing.T) {
	// START 3 OMIT
	type Stats struct {
		sync.Mutex
		Reads  int64
		Writes int64
	}
	// END 3 OMIT

	// START 4 OMIT
	s := Stats{}
	s.Lock()
	defer s.Unlock()
	// END 4 OMIT
}

// START 5 OMIT
func TestRaceCounter(t *testing.T) {
	counter := new(int32) // HLxxx
	var wg sync.WaitGroup
	// multiple producers
	for producerId := 0; producerId < 10; producerId++ {
		wg.Add(1)

		// spawn a new counter
		go func(producerId int) {
			defer wg.Done()
			atomic.AddInt32(counter, 1) // HLxxx
		}(producerId)
	}
	wg.Wait()                            // wait for the N producers
	if atomic.LoadInt32(counter) != 10 { // HLxxx
		t.Fatal("counter failed.")
	}
}

// END 5 OMIT

// START 6 OMIT
func TestRaceCounterWithChannels(t *testing.T) {
	counter := 0                 // HLxxx
	increments := make(chan int) // HLxxx
	var wg sync.WaitGroup
	for producerId := 0; producerId < 10; producerId++ {
		wg.Add(1)
		// spawn a new counter
		go func(producerId int, increments chan int) {
			defer wg.Done()
			increments <- 1 // HLxxx
		}(producerId, increments)
	}
	go func() {
		wg.Wait()
		close(increments)
	}()
	for increment := range increments { // HLxxx
		counter = counter + increment // HLxxx
	} // HLxxx
	if counter != 10 {
		t.Fatal("counter failed.")
	}
}

// END 6 OMIT


// START 7 OMIT
func TestRaceCounterWaitGroups(t *testing.T) {
	var wg sync.WaitGroup{} // HLxxx
	for n := 0; n < 100; n++ {
		wg.Add(1) // HLxxx
		go func() {
			defer wg.Done() // HLxxx

			HeavyTask()
		}()
	}
	wg.Wait() // HLxxx
}
// END 7 OMIT

// START 8 OMIT
func TestRaceCounterErrGroups(t *testing.T) {
	var eg errgroup.Group{}  // HLxxx

	for n := 0; n < 100; n++ {
		eg.Go(func() error {  // HLxxx
			return nil
		})
	}
	if err := eg.Wait(); err != nil {  // HLxxx
		t.Fatal(err)
	}
}
// END 8 OMIT