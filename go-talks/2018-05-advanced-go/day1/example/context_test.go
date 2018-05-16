package context_test

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

func SomeTask(v struct{}) {

}

func TestContextTest(t *testing.T) {

	ch := make(chan struct{})

	// START 1 OMIT
	ctx, cancel := context.WithCancel(context.Background()) // HLxxx
	go func() {
		select {
		case <-ctx.Done(): // HLxxx
			fmt.Printf("canceled!\n")
		case value := <-ch:
			SomeTask(value)
		}
	}()
	cancel() // cancels the operation // HLxxx

	// END 1 OMIT
}

func TestContextWithDeadline(t *testing.T) {

	ch := make(chan struct{})

	var wg sync.WaitGroup
	wg.Add(1)

	// START 2 OMIT
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(30*time.Second)) // HLxxx

	go func() {
		defer wg.Done()

		select {
		case <-ctx.Done(): // HLxxx
			fmt.Printf("deadline!\n")
		case value := <-ch:
			SomeTask(value)
		}
	}()

	wg.Wait()
	// END 2 OMIT

	_ = cancel
}

// START 3 OMIT
func SomeWork(ctx context.Context) { // HLxxx
	k := ctx.Value("some-value") // HLxxx
	if value, ok := k.(int); !ok || value != 1492 {
		panic("invalid 'some-value' context-value")
	}
	k := ctx.Value("some-other-value") // HLxxx
	if k != nil {
		panic("'some-other-value' should be nil")
	}
}
func TestContextWithDeadline(t *testing.T) {
	ctx := context.WithValue(context.Background(), "some-value", 1492) // HLxxx
	SomeWork(ctx)
}

// END 3 OMIT
