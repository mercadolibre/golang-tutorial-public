package main

import (
	"sync"
)

type Item struct{}

func ProcessItem(Item) {}

func pattern1() {
	numberOfConsumers := 10
	numberOfProducers := 10
	items := make(chan Item)
	shouldProduce := true

	// START 1 OMIT
	// multiple-consumers
	for consumerId := 0; consumerId < numberOfConsumers; consumerId++ {
		go func(consumerId int, items chan Item) {
			for item := range items { // HLxxx
				ProcessItem(item)
			} // HLxxx
		}(consumerId, items)
	}

	// single-producer
	for shouldProduce {
		items <- Item{} // HLxxx
	}
	close(items) // HLxxx
	// END 1 OMIT

	// START 2 OMIT
	// multiple producers
	for producerId := 0; producerId < numberOfProducers; producerId++ {
		go func(producerId int, items chan Item) {
			for shouldProduce {
				items <- Item{} // HLxxx
			}
		}(producerId, items)
	}

	// single consumer
	for item := range items { // HLxxx
		ProcessItem(item)
	}
	// END 2 OMIT

	// START 3 OMIT
	// multiple producers
	var wg sync.WaitGroup // HLxxx
	for producerId := 0; producerId < numberOfProducers; producerId++ {
		wg.Add(1) // HLxxx
		go func(producerId int, items chan Item) {
			defer wg.Done() // HLxxx
			for shouldProduce {
				items <- Item{}
			}
		}(producerId, items)
	}
	go func() { // HLxxx
		wg.Wait()    // HLxxx
		close(items) // HLxxx
	}() // HLxxx

	// single consumer
	for item := range items {
		ProcessItem(item)
	}
	// END 3 OMIT
}
