package main

import (
	"fmt"
	"strings"
	"sync"
)

// START 1 OMIT
func main() {
	emojis := []string{"🤷", "❤", "♡", "😂", "🤔", "😍", "🔥", "💩", "😊", "😘"}

	ch := make(chan string)
	defer close(ch)

	var wg sync.WaitGroup

	go func() {
		for e := range ch {
			fmt.Println(e)
		}
	}()

	for i, e := range emojis {
		wg.Add(1)
		go func(e string, i int) {
			defer wg.Done()
			ch <- strings.Repeat(e, i+1)
		}(e, i)
	}

	wg.Wait()
}

// END 1 OMIT
