package main
import (
	"sync"
	"time"
	"fmt"
)
func main() {
	var myWaitGroup sync.WaitGroup
	myWaitGroup.Add(2)
	go func() {
		fmt.Println("Start goroutine 1")
		fmt.Println("End goroutine 1")
		myWaitGroup.Done()
	}()
	go func() {
		fmt.Println("Start goroutine 2")
		time.Sleep(time.Second * 5)
		fmt.Println("End goroutine 2")
		myWaitGroup.Done()
	}()
	fmt.Println("Waiting for all goroutines to exit")
	myWaitGroup.Wait()
	fmt.Println("Waited for all goroutines to exit")
}