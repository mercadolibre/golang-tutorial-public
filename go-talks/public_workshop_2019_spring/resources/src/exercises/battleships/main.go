package main

import (
	"os"
	"os/signal"

	tm "github.com/buger/goterm"
)

func main() {
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt)
	go signalHandler(signalChannel)

	engine := NewEngine([]PlayerRunnerFunction{
		RandomAI,
		RandomAI,
	})

	engine.run()
}

func signalHandler(signalChannel chan os.Signal) {
	for _ = range signalChannel {
		tm.Clear()
		tm.MoveCursor(1, 1)
		tm.Flush()
		os.Exit(1)
	}
}
